package service

import (
	"ae86/internal/model"
	"ae86/internal/service/adapter"
	"errors"
	"gorm.io/gorm"
)

type CustomerService struct {
	storage adapter.StorageContainer
}

func NewCustomerService(storage adapter.StorageContainer) *CustomerService {
	return &CustomerService{storage: storage}
}

func (c *CustomerService) ExistsByExternalID(externalID uint) (result bool, err error) {
	customer, err := c.storage.Customer().ByExternalID(externalID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	result = customer.ID != 0
	return
}

func (c *CustomerService) ByExternalID(externalID uint) (result model.Customer, err error) {
	return c.storage.Customer().ByExternalID(externalID)
}

func (c *CustomerService) Create(customer model.Customer) (id uint, err error) {
	return c.storage.Customer().Create(customer)
}
