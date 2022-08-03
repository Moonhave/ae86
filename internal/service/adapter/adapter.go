package adapter

import "ae86/internal/model"

type StorageContainer interface {
	Category() CategoryStorage
	Product() ProductStorage
	Customer() CustomerStorage
	Order() OrderStorage
	OrderItem() OrderItemStorage
}

type CategoryStorage interface {
	GetByID(id uint) (result model.Category, err error)
	Create(category model.Category) (id uint, err error)
	Update(id uint, category model.Category) (err error)
	Delete(id uint) (err error)
}

type ProductStorage interface {
	GetByID(id uint) (result model.Product, err error)
	GetAllBy(filter ProductFilter) (result []model.Product, err error)
	Create(product model.Product) (id uint, err error)
	Update(id uint, product model.Product) (err error)
	Delete(id uint) (err error)
}

type CustomerStorage interface {
	GetByID(id uint) (result model.Customer, err error)
	GetByExternalID(externalID uint) (result model.Customer, err error)
	Create(customer model.Customer) (id uint, err error)
}

type OrderStorage interface {
	GetByID(id uint) (result model.Order, err error)
	GetAllBy(filter OrderFilter) (result []model.Order, err error)
	Create(order model.Order) (id uint, err error)
	Update(id uint, order model.Order) (err error)
	Delete(id uint) (err error)
}

type OrderItemStorage interface {
	GetByID(id uint) (result model.OrderItem, err error)
	GetAllByOrderID(orderID uint) (result []model.OrderItem, err error)
	Create(orderItem model.OrderItem) (id uint, err error)
	Update(id uint, orderItem model.OrderItem) (err error)
	Delete(id uint) (err error)
}
