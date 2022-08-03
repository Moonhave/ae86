package storage

import (
	"ae86/internal/model"
	"ae86/internal/service/adapter"
	"ae86/pkg/logger"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type OrderStorage struct {
	db *gorm.DB
}

func NewOrderStorage(db *gorm.DB) *OrderStorage {
	return &OrderStorage{db: db}
}

func (s *OrderStorage) ByID(id uint) (result model.Order, err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error": err,
				"id":    id,
			}).Error("OrderStorage.GetByID failed")
		}
	}()

	err = s.db.
		Model(&model.Order{}).
		Where("id = ?", id).
		First(&result).
		Error
	return
}

func (s *OrderStorage) ListBy(filter adapter.OrderFilter) (result []model.Order, err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error":  err,
				"filter": filter,
			}).Error("OrderStorage.GetAllBy failed")
		}
	}()

	err = s.db.
		Model(&model.Order{}).
		Scopes(func(db *gorm.DB) *gorm.DB {
			if filter.Address != nil && *filter.Address != "" {
				db = db.Where("LOWER(address) LIKE LOWER(?)", "%"+(*filter.Address)+"%")
			}
			if filter.State != nil && *filter.State != "" {
				db = db.Where("state = ?", filter.State)
			}
			if filter.PaymentMethod != nil && *filter.PaymentMethod != "" {
				db = db.Where("payment_method = ?", filter.PaymentMethod)
			}
			if filter.CustomerID != nil && *filter.CustomerID != 0 {
				db = db.Where("customer_id = ?", filter.CustomerID)
			}
			return db
		}).
		Find(&result).
		Error
	return
}

func (s *OrderStorage) Create(order model.Order) (id uint, err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error": err,
				"order": order,
			}).Error("OrderStorage.Create failed")
		}
	}()

	err = s.db.
		Model(&model.Order{}).
		Create(&order).
		Error
	id = order.ID
	return
}

func (s *OrderStorage) Update(id uint, order model.Order) (err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error": err,
				"id":    id,
				"order": order,
			}).Error("OrderStorage.Update failed")
		}
	}()

	err = s.db.
		Model(&model.Order{}).
		Where("id = ?", id).
		Updates(&order).
		Error
	return
}

func (s *OrderStorage) Delete(id uint) (err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error": err,
				"id":    id,
			}).Error("OrderStorage.Delete failed")
		}
	}()

	err = s.db.
		Model(&model.Order{}).
		Where("id = ?", id).
		Delete(&model.Order{}).
		Error
	return
}
