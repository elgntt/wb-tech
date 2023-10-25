package db

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"wb-tech/internal/config"
)

func OpenDB(ctx context.Context, cfg config.DBConfig) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PgHost, cfg.PgPort, cfg.PgUser, cfg.PgPassword, cfg.PgDatabase,
	)
	log.Println(connectionString)
	db, err := sqlx.Open("pgx", connectionString)
	if err != nil {
		return nil, fmt.Errorf("OpenDB sqlx open: %w", err)
	}

	if err = db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("OpenDB sqlx ping: %w", err)
	}

	return db, nil
}
