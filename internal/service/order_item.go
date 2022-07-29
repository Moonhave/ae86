package service

import (
	"ae86/internal/service/adapter"
)

type OrderItemService struct {
	storage adapter.StorageContainer
}

func NewOrderItemService(storage adapter.StorageContainer) *OrderItemService {
	return &OrderItemService{storage: storage}
}
