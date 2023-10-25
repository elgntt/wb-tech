package service

import (
	"github.com/nats-io/stan.go"
	"log"
	"wb-tech/internal/model"
	"wb-tech/internal/pkg/app_err"
)

import (
	"context"
	"encoding/json"
)

type Service struct {
	sc    stan.Conn
	repo  repository
	Cache map[string]model.OrderData
}

func New(sc stan.Conn, repo repository) *Service {
	return &Service{
		sc:    sc,
		repo:  repo,
		Cache: make(map[string]model.OrderData),
	}
}

func (s *Service) GetFromCache(uid string) (model.OrderData, error) {
	orderData, ok := s.Cache[uid]
	if !ok {
		return model.OrderData{}, app_err.NewBusinessError("order not found!")
	}

	return orderData, nil
}

func (s *Service) LoadCache(ctx context.Context) error {
	cacheDB, err := s.repo.LoadCache(ctx)
	if err != nil {
		return err
	}

	s.Cache = cacheDB

	return nil
}

func (s *Service) Listen(ctx context.Context, listenChannel string) error {
	var err error

	sub, err := s.sc.Subscribe(listenChannel, func(m *stan.Msg) {
		data := model.OrderData{}
		err := json.Unmarshal(m.Data, &data)
		if err != nil {
			log.Printf("NATS: unmarshal: %v", err)
			return
		}

		err = data.Validate()
		if err != nil {
			log.Println(err)
			return
		}

		for i := range data.Items {
			data.Items[i].OrderUid = data.OrderUid
		}

		err = s.repo.Save(ctx, data)
		if err != nil {
			log.Println(err)
			return
		}

		s.Cache[data.OrderUid] = data
	})
	if err != nil {
		return err
	}

	defer func() {
		err := sub.Unsubscribe()
		if err != nil {
			log.Println(err)
		}
		// Закрываем соединение
		err = s.sc.Close()
		if err != nil {
			log.Println(err)
		}

	}()

	<-ctx.Done()

	return nil
}
