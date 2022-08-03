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

func (p *ProductService) ListByCategoryID(categoryID uint) (result []model.Product, err error) {
	return p.storage.Product().ListBy(adapter.ProductFilter{CategoryID: &categoryID})
}
