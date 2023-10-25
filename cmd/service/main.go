package main

import (
	"context"
	"errors"
	"github.com/nats-io/stan.go"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"wb-tech/internal/api"
	"wb-tech/internal/config"
	db2 "wb-tech/internal/pkg/db"
	"wb-tech/internal/repository"
	service2 "wb-tech/internal/service"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	sc, err := stan.Connect(cfg.NatsStreamingConfig.ClusterId, cfg.NatsStreamingConfig.ClientId, stan.NatsURL(cfg.NatsStreamingConfig.ListenUrl))
	if err != nil {
		log.Fatalf("Failed to connect to NATS Streaming: %v", err)
	}

	db, err := db2.OpenDB(ctx, cfg.DBConfig)
	if err != nil {
		log.Fatalf("Failed to open the database: %v", err)
	}

	service := service2.New(
		sc,
		repository.New(db),
	)
	err = service.LoadCache(ctx)
	if err != nil {
		log.Fatalf("Failed to load cache from the database: %v", err)
	}

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return service.Listen(ctx, cfg.NatsStreamingConfig.ListenChannel)
	})

	r := api.New(service)

	srv := &http.Server{
		Addr:    ":" + cfg.ServerConfig.HTTPPort,
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Failed to listen: %v", err)
		}
	}()

	g.Go(func() error {
		<-gCtx.Done()
		return srv.Shutdown(ctx)
	})

	if err := g.Wait(); err != nil {
		log.Printf("exit reason: %s \n", err)
	}
}
