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
	ListAll() (result []model.Category, err error)
	ByID(id uint) (result model.Category, err error)
	Create(category model.Category) (id uint, err error)
	Update(id uint, category model.Category) (err error)
	Delete(id uint) (err error)
}

type ProductStorage interface {
	ByID(id uint) (result model.Product, err error)
	ListBy(filter ProductFilter) (result []model.Product, err error)
	Create(product model.Product) (id uint, err error)
	Update(id uint, product model.Product) (err error)
	Delete(id uint) (err error)
}

type CustomerStorage interface {
	ByID(id uint) (result model.Customer, err error)
	ByExternalID(externalID uint) (result model.Customer, err error)
	IsExistsByExternalID(externalID uint) (result bool, err error)
	Create(customer model.Customer) (id uint, err error)
}

type OrderStorage interface {
	ByID(id uint) (result model.Order, err error)
	ListBy(filter OrderFilter) (result []model.Order, err error)
	Create(order model.Order) (id uint, err error)
	Update(id uint, order model.Order) (err error)
	Delete(id uint) (err error)
}

type OrderItemStorage interface {
	ByID(id uint) (result model.OrderItem, err error)
	ListByOrderID(orderID uint) (result []model.OrderItem, err error)
	Create(orderItem model.OrderItem) (id uint, err error)
	Update(id uint, orderItem model.OrderItem) (err error)
	Delete(id uint) (err error)
}
