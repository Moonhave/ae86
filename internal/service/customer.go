package service

import (
	"ae86/internal/model"
	"ae86/internal/service/adapter"
)

type CustomerService struct {
	storage adapter.StorageContainer
}

func (c CustomerService) CreateCustomer(customer model.Customer) (id uint, err error) {
	return 0, nil
}

func NewCustomerService(storage adapter.StorageContainer) *CustomerService {
	return &CustomerService{storage: storage}
}
