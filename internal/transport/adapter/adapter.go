package adapter

import "ae86/internal/model"

type ServiceContainer interface {
	Manager() ManagerService
	Store() StoreService
	Category() CategoryService
	Product() ProductService
	Customer() CustomerService
	Order() OrderService
	OrderItem() OrderItemService
}

type ManagerService interface {
	GetAllManagers() (result []model.Manager)
}

type CategoryService interface {
	GetAllCategories() (result []model.Category)
}

type StoreService interface{}

type ProductService interface {
	GetProductByCategory(categoryId int) (result []model.Product)
}

type CustomerService interface {
	CreateCustomer(customer model.Customer) (id uint, err error)
}

type OrderService interface {
	GetOrderList(filter OrderFilter) (result []model.Order, err error)
}

type OrderItemService interface {
	CreateOrder(order model.Order) (id uint, err error)
}
