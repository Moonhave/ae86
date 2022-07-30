package service

import (
	"ae86/internal/model"
	"ae86/internal/service/adapter"
	transportAdapter "ae86/internal/transport/adapter"
)

type OrderService struct {
	storage adapter.StorageContainer
}

func (o *OrderService) GetOrderList(filter transportAdapter.OrderFilter) (result []model.Order, err error) {
	orderFilter := adapter.OrderFilter{
		Address:       filter.Address,
		CustomerID:    filter.CustomerID,
		PaymentMethod: filter.PaymentMethod,
		State:         filter.State,
		StoreID:       filter.StoreID,
		IsDeleted:     filter.IsDeleted,
	}

	result, err = o.storage.Order().GetAllBy(orderFilter)
	if result != nil {
		return nil, err
	}
	return result, err
}

func (o *OrderService) CreateOrder(order model.Order) (id uint, err error) {
	result, err := o.storage.Order().Create(order)
	if err != nil {
		return 0, err
	}
	return result, err
}

func NewOrderService(storage adapter.StorageContainer) *OrderService {
	return &OrderService{storage: storage}
}
