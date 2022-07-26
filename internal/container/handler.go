package container

import (
	"ae86/internal/transport/adapter"
	"ae86/internal/transport/bot/handlers"
)

type HandlerContainer struct {
	category *handlers.CategoryHandler
}

func NewHandlerContainer(serviceContainer adapter.ServiceContainer) *HandlerContainer {
	return &HandlerContainer{
		category: handlers.NewCategoryHandler(serviceContainer),
	}
}

func (c *HandlerContainer) Category() *handlers.CategoryHandler {
	return c.category
}
