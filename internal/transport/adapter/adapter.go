package adapter

import "ae86/internal/model"

type ServiceContainer interface {
	Category() CategoryService
	Product() ProductService
	Customer() CustomerService
	Order() OrderService
	OrderItem() OrderItemService
}

type CategoryService interface {
	ListAll() (result []model.Category, err error)
}

type ProductService interface {
	ListByCategoryID(categoryID uint) (result []model.Product, err error)
}

type CustomerService interface {
	Create(customer model.Customer) (id uint, err error)
}

type OrderService interface {
	Create(order model.Order) (id uint, err error)
	ListBy(filter OrderFilter) (result []model.Order, err error)
}

type OrderItemService interface{}
