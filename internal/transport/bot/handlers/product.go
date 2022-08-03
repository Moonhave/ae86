package handlers

import (
	"ae86/internal/model"
	"ae86/internal/transport/adapter"
	"ae86/internal/transport/bot/view"
	tele "gopkg.in/telebot.v3"
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

func (h *ProductHandler) UpdateProductInlineMenu(c tele.Context) error {
	buttonRows := []tele.Row{
		view.ProductMenu.Row(view.BtnInlineAdded),
		view.ProductMenu.Row(view.BtnInlineOrder),
		view.ProductMenu.Row(view.BtnInlineBack),
	}
	view.ProductMenu.Inline(buttonRows...)
	c.Edit(view.ProductMenu)

	return c.Respond(&tele.CallbackResponse{Text: view.AddedToCartMessage})
}
