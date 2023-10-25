package api

import "wb-tech/internal/model"

type Service interface {
	GetFromCache(uid string) (model.OrderData, error)
}
