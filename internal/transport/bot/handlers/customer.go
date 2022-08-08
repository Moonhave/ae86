package handlers

import (
	"ae86/internal/transport/adapter"
)

func NewCustomerHandler(service adapter.ServiceContainer) *CustomerHandler {
	return &CustomerHandler{service: service}
}

type CustomerHandler struct {
	service adapter.ServiceContainer
}
