package service

import (
	"context"
	"wb-tech/internal/model"
)

type repository interface {
	Save(ctx context.Context, orderData model.OrderData) (err error)
	LoadCache(ctx context.Context) (map[string]model.OrderData, error)
}
