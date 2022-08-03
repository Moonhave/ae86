package storage

import (
	"ae86/internal/model"
	"ae86/internal/service/adapter"
	"ae86/pkg/logger"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductStorage struct {
	db *gorm.DB
}

func NewProductStorage(db *gorm.DB) *ProductStorage {
	return &ProductStorage{db: db}
}

func (s *ProductStorage) GetByID(id uint) (result model.Product, err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error": err,
				"id":    id,
			}).Error("ProductStorage.GetByID failed")
		}
	}()

	err = s.db.
		Model(&model.Product{}).
		Where("id = ?", id).
		First(&result).
		Error
	return
}

func (s *ProductStorage) GetAllBy(filter adapter.ProductFilter) (result []model.Product, err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error":  err,
				"filter": filter,
			}).Error("ProductStorage.GetAllBy failed")
		}
	}()

	err = s.db.
		Model(&model.Product{}).
		Scopes(func(db *gorm.DB) *gorm.DB {
			if filter.Title != nil && *filter.Title != "" {
				db = db.Where("LOWER(title) LIKE LOWER(?)", "%"+(*filter.Title)+"%")
			}
			if filter.MinPrice != nil && *filter.MinPrice != 0 {
				db = db.Where("price >= ?", filter.MinPrice)
			}
			if filter.MaxPrice != nil && *filter.MaxPrice != 0 {
				db = db.Where("price <= ?", filter.MaxPrice)
			}
			if filter.IsActive != nil {
				db = db.Where("is_active = ?", filter.IsActive)
			}
			if filter.CategoryID != nil && *filter.CategoryID != 0 {
				db = db.Where("category_id = ?", filter.CategoryID)
			}
			return db
		}).
		Find(&result).
		Error
	return
}

func (s *ProductStorage) Create(product model.Product) (id uint, err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error":   err,
				"product": product,
			}).Error("ProductStorage.Create failed")
		}
	}()

	err = s.db.
		Model(&model.Product{}).
		Create(&product).
		Error
	id = product.ID
	return
}

func (s *ProductStorage) Update(id uint, product model.Product) (err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error":   err,
				"id":      id,
				"product": product,
			}).Error("ProductStorage.Update failed")
		}
	}()

	err = s.db.
		Model(&model.Product{}).
		Where("id = ?", id).
		Updates(&product).
		Error
	return
}

func (s *ProductStorage) Delete(id uint) (err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error": err,
				"id":    id,
			}).Error("ProductStorage.Delete failed")
		}
	}()

	err = s.db.
		Model(&model.Product{}).
		Where("id = ?", id).
		Delete(&model.Product{}).
		Error
	return
}
