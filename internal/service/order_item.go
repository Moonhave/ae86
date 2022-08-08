package service

import (
	"ae86/internal/model"
	"ae86/internal/service/adapter"
)

type OrderItemService struct {
	storage adapter.StorageContainer
}

func NewOrderItemService(storage adapter.StorageContainer) *OrderItemService {
	return &OrderItemService{storage: storage}
}

func (o *OrderItemService) Create(item model.OrderItem) (uint, error) {
	return o.storage.OrderItem().Create(item)
}

func (o *OrderItemService) ListByOrderID(id uint) ([]model.OrderItem, error) {
	return o.storage.OrderItem().ListByOrderID(id)
}
