package service

import (
	"ae86/internal/model"
	"ae86/internal/service/adapter"
	transportAdapter "ae86/internal/transport/adapter"
)

type OrderService struct {
	storage adapter.StorageContainer
}

func (o OrderService) GetOrderList(filter transportAdapter.OrderFilter) (result []model.Order, err error) {
	return nil, nil
}

func (o OrderService) CreateOrder(order model.Order) (id uint, err error) {
	return 0, nil
}

func NewOrderService(storage adapter.StorageContainer) *OrderService {
	return &OrderService{storage: storage}
}
