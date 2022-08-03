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
	GetAllCategories() (result []model.Category, err error)
}

type ProductService interface {
	GetProductsByCategory(categoryId uint) (result []model.Product, err error)
}

type CustomerService interface {
	CreateCustomer(customer model.Customer) (id uint, err error)
}

type OrderService interface {
	CreateOrder(order model.Order) (id uint, err error)
	GetOrderList(filter OrderFilter) (result []model.Order, err error)
}

type OrderItemService interface{}
