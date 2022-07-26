package service

import (
	"ae86/internal/model"
	"ae86/internal/service/adapter"
)

type CategoryService struct {
	storage adapter.StorageContainer
}

func NewCategoryService(storage adapter.StorageContainer) *CategoryService {
	return &CategoryService{storage: storage}
}

func (c *CategoryService) ListAll() (result []model.Category, err error) {
	return c.storage.Category().ListAll()
}
