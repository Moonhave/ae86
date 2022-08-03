package container

import (
	"ae86/internal/service/adapter"
	"ae86/internal/storage"
	"gorm.io/gorm"
)

type storageContainer struct {
	category  *storage.CategoryStorage
	product   *storage.ProductStorage
	customer  *storage.CustomerStorage
	order     *storage.OrderStorage
	orderItem *storage.OrderItemStorage
}

func NewStorageContainer(db *gorm.DB) *storageContainer {
	return &storageContainer{
		category:  storage.NewCategoryStorage(db),
		product:   storage.NewProductStorage(db),
		customer:  storage.NewCustomerStorage(db),
		order:     storage.NewOrderStorage(db),
		orderItem: storage.NewOrderItemStorage(db),
	}
}

func (c *storageContainer) Category() adapter.CategoryStorage {
	return c.category
}

func (c *storageContainer) Product() adapter.ProductStorage {
	return c.product
}

func (c *storageContainer) Customer() adapter.CustomerStorage {
	return c.customer
}

func (c *storageContainer) Order() adapter.OrderStorage {
	return c.order
}

func (c *storageContainer) OrderItem() adapter.OrderItemStorage {
	return c.orderItem
}
