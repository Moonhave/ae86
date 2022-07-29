package service

import (
	"ae86/internal/model"
	"ae86/internal/service/adapter"
	"log"
)

type CategoryService struct {
	storage adapter.StorageContainer
}

func NewCategoryService(storage adapter.StorageContainer) *CategoryService {
	return &CategoryService{storage: storage}
}

func (c *CategoryService) GetAllCategories() (result []model.Category) {
	result, err := c.storage.Category().GetAllByStoreID(0)
	if err != nil {
		log.Println(err)
	}
	return result
}
