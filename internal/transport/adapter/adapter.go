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
	ByID(id uint) (result model.Product, err error)
	ListByCategoryID(categoryID uint) (result []model.Product, err error)
}

type CustomerService interface {
	ExistsByExternalID(externalID uint) (result bool, err error)
	ByExternalID(externalID uint) (result model.Customer, err error)
	Create(customer model.Customer) (id uint, err error)
}

type OrderService interface {
	Create(order model.Order) (id uint, err error)
	ListBy(filter OrderFilter) (result []model.Order, err error)
	Update(id uint, order model.Order) (err error)
	Delete(id uint) (err error)
}

type OrderItemService interface {
	Create(orderItem model.OrderItem) (id uint, err error)
	ListByOrderID(orderID uint) (result []model.OrderItem, err error)
}
