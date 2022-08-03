package container

import (
	"ae86/internal/service"
	serviceAdapter "ae86/internal/service/adapter"
	transportAdapter "ae86/internal/transport/adapter"
)

type serviceContainer struct {
	category  *service.CategoryService
	product   *service.ProductService
	customer  *service.CustomerService
	order     *service.OrderService
	orderItem *service.OrderItemService
}

func NewServiceContainer(storageContainer serviceAdapter.StorageContainer) *serviceContainer {
	return &serviceContainer{
		category:  service.NewCategoryService(storageContainer),
		product:   service.NewProductService(storageContainer),
		customer:  service.NewCustomerService(storageContainer),
		order:     service.NewOrderService(storageContainer),
		orderItem: service.NewOrderItemService(storageContainer),
	}
}

func (c *serviceContainer) Category() transportAdapter.CategoryService {
	return c.category
}

func (c *serviceContainer) Product() transportAdapter.ProductService {
	return c.product
}

func (c *serviceContainer) Customer() transportAdapter.CustomerService {
	return c.customer
}

func (c *serviceContainer) Order() transportAdapter.OrderService {
	return c.order
}

func (c *serviceContainer) OrderItem() transportAdapter.OrderItemService {
	return c.orderItem
}
