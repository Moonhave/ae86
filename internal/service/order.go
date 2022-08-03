package service

import (
	"ae86/internal/model"
	"ae86/internal/service/adapter"
	transportAdapter "ae86/internal/transport/adapter"
)

type OrderService struct {
	storage adapter.StorageContainer
}

func NewOrderService(storage adapter.StorageContainer) *OrderService {
	return &OrderService{storage: storage}
}

func (o *OrderService) ListBy(filter transportAdapter.OrderFilter) (result []model.Order, err error) {
	orderFilter := adapter.OrderFilter{
		Address:       filter.Address,
		CustomerID:    filter.CustomerID,
		PaymentMethod: filter.PaymentMethod,
		State:         filter.State,
	}

	return o.storage.Order().ListBy(orderFilter)
}

func (o *OrderService) Create(order model.Order) (id uint, err error) {
	return o.storage.Order().Create(order)
}
