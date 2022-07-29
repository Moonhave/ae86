package service

import (
	"ae86/internal/model"
	"ae86/internal/service/adapter"
)

type ProductService struct {
	storage adapter.StorageContainer
}

func (p ProductService) GetProductByCategory(categoryId int) (result []model.Product) {
	return nil
}

func NewProductService(storage adapter.StorageContainer) *ProductService {
	return &ProductService{storage: storage}
}
