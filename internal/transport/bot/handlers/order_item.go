package handlers

import (
	"ae86/internal/transport/adapter"
)

func NewOrderItemHandler(service adapter.ServiceContainer) *OrderItemHandler {
	return &OrderItemHandler{service: service}
}

type OrderItemHandler struct {
	service adapter.ServiceContainer
}
