package service

import (
	"ae86/internal/model"
	"ae86/internal/service/adapter"
)

type CustomerService struct {
	storage adapter.StorageContainer
}

func NewCustomerService(storage adapter.StorageContainer) *CustomerService {
	return &CustomerService{storage: storage}
}

func (c *CustomerService) CreateCustomer(customer model.Customer) (id uint, err error) {
	result, err := c.storage.Customer().Create(customer)
	if err != nil {
		return 0, err
	}
	return result, err
}
