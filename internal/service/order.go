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

func (o *OrderService) Update(id uint, order model.Order) (err error) {
	return o.storage.Order().Update(id, order)
}

func (o *OrderService) Delete(id uint) (err error) {
	orderItems, err := o.storage.OrderItem().ListByOrderID(id)
	if err != nil {
		return err
	}

	for _, orderItem := range orderItems {
		err = o.storage.OrderItem().Delete(orderItem.ID)
		if err != nil {
			return err
		}
	}

	return o.storage.Order().Delete(id)
}
