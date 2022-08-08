package handlers

import (
	"ae86/internal/model"
	"ae86/internal/transport/adapter"
)

func NewProductHandler(service adapter.ServiceContainer) *ProductHandler {
	return &ProductHandler{service: service}
}

type ProductHandler struct {
	service adapter.ServiceContainer
}

func (h *ProductHandler) GetAllProductsByCategoryID(categoryId uint) ([]model.Product, error) {
	products, err := h.service.Product().ListByCategoryID(categoryId)
	if err != nil {
		return []model.Product{}, err
	}
	return products, err
}
