package service

import (
	"ae86/internal/model"
	"ae86/internal/service/adapter"
)

type ProductService struct {
	storage adapter.StorageContainer
}

func NewProductService(storage adapter.StorageContainer) *ProductService {
	return &ProductService{storage: storage}
}

func (p ProductService) GetProductByCategory(categoryId int) (result []model.Product) {
	return nil
}

func (p *ProductService) GetProductsByCategory(categoryId uint) (result []model.Product, err error) {
	return p.storage.Product().GetAllBy(adapter.ProductFilter{CategoryID: &categoryId})
}
