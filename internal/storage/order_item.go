package storage

import (
	"ae86/internal/model"
	"ae86/pkg/logger"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type OrderItemStorage struct {
	db *gorm.DB
}

func NewOrderItemStorage(db *gorm.DB) *OrderItemStorage {
	return &OrderItemStorage{db: db}
}

func (s *OrderItemStorage) ByID(id uint) (result model.OrderItem, err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error": err,
				"id":    id,
			}).Error("OrderItemStorage.GetByID failed")
		}
	}()

	err = s.db.
		Model(&model.OrderItem{}).
		Where("id = ?", id).
		First(&result).
		Error
	return
}

func (s *OrderItemStorage) ListByOrderID(orderID uint) (result []model.OrderItem, err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error":   err,
				"orderID": orderID,
			}).Error("OrderItemStorage.GetAllByOrderID failed")
		}
	}()

	err = s.db.
		Model(&model.OrderItem{}).
		Where("order_id = ?", orderID).
		Find(&result).
		Error
	return
}

func (s *OrderItemStorage) Create(orderItem model.OrderItem) (id uint, err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error":     err,
				"orderItem": orderItem,
			}).Error("OrderItemStorage.Create failed")
		}
	}()

	err = s.db.
		Model(&model.OrderItem{}).
		Create(&orderItem).
		Error
	id = orderItem.ID
	return
}

func (s *OrderItemStorage) Update(id uint, orderItem model.OrderItem) (err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error":     err,
				"id":        id,
				"orderItem": orderItem,
			}).Error("OrderItemStorage.Update failed")
		}
	}()

	err = s.db.
		Model(&model.OrderItem{}).
		Where("id = ?", id).
		Updates(&orderItem).
		Error
	return
}

func (s *OrderItemStorage) Delete(id uint) (err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error": err,
				"id":    id,
			}).Error("OrderItemStorage.Delete failed")
		}
	}()

	err = s.db.
		Model(&model.OrderItem{}).
		Where("id = ?", id).
		Delete(&model.OrderItem{}).
		Error
	return
}
