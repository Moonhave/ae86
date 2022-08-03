package container

import (
	"ae86/internal/transport/adapter"
	"ae86/internal/transport/bot/handlers"
)

type HandlerContainer struct {
	general   *handlers.GeneralHandler
	category  *handlers.CategoryHandler
	customer  *handlers.CustomerHandler
	order     *handlers.OrderHandler
	orderItem *handlers.OrderItemHandler
	product   *handlers.ProductHandler
}

func NewHandlerContainer(serviceContainer adapter.ServiceContainer) *HandlerContainer {
	return &HandlerContainer{
		general:   handlers.NewGeneralHandler(serviceContainer),
		category:  handlers.NewCategoryHandler(serviceContainer),
		customer:  handlers.NewCustomerHandler(serviceContainer),
		order:     handlers.NewOrderHandler(serviceContainer),
		orderItem: handlers.NewOrderItemHandler(serviceContainer),
		product:   handlers.NewProductHandler(serviceContainer),
	}
}

func (c *HandlerContainer) General() *handlers.GeneralHandler {
	return c.general
}

func (c *HandlerContainer) Category() *handlers.CategoryHandler {
	return c.category
}

func (c *HandlerContainer) Customer() *handlers.CustomerHandler {
	return c.customer
}

func (c *HandlerContainer) Order() *handlers.OrderHandler {
	return c.order
}

func (c *HandlerContainer) OrderItem() *handlers.OrderItemHandler {
	return c.orderItem
}

func (c *HandlerContainer) Product() *handlers.ProductHandler {
	return c.product
}
