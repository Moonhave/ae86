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

func (c *CustomerService) ExistsByExternalID(externalID uint) (result bool, err error) {
	return c.storage.Customer().IsExistsByExternalID(externalID)
}

func (c *CustomerService) ByExternalID(externalID uint) (result model.Customer, err error) {
	return c.storage.Customer().ByExternalID(externalID)
}

func (c *CustomerService) Create(customer model.Customer) (id uint, err error) {
	return c.storage.Customer().Create(customer)
}

func (c *CustomerService) Update(id uint, customer model.Customer) (err error) {
	return c.storage.Customer().Update(id, customer)
}
