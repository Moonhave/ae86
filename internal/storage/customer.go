package storage

import (
	"ae86/internal/model"
	"ae86/pkg/logger"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CustomerStorage struct {
	db *gorm.DB
}

func NewCustomerStorage(db *gorm.DB) *CustomerStorage {
	return &CustomerStorage{db: db}
}

func (s *CustomerStorage) ByID(id uint) (result model.Customer, err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error": err,
				"id":    id,
			}).Error("CustomerStorage.GetByID failed")
		}
	}()

	err = s.db.
		Model(&model.Customer{}).
		Where("id = ?", id).
		First(&result).
		Error
	return
}

func (s *CustomerStorage) ByExternalID(externalID uint) (result model.Customer, err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error":      err,
				"externalID": externalID,
			}).Error("CustomerStorage.GetByExternalID failed")
		}
	}()

	err = s.db.
		Model(&model.Customer{}).
		Where("external_id = ?", externalID).
		First(&result).
		Error
	return
}

func (s *CustomerStorage) IsExistsByExternalID(externalID uint) (result bool, err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error":      err,
				"externalID": externalID,
			}).Error("CustomerStorage.IsExistsByExternalID failed")
		}
	}()

	c := model.Customer{}
	err = s.db.
		Model(&model.Customer{}).
		Where("external_id = ?", externalID).
		First(&c).
		Error
	result = c.ID != 0
	return
}

func (s *CustomerStorage) Create(customer model.Customer) (id uint, err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error":    err,
				"customer": customer,
			}).Error("CustomerStorage.Create failed")
		}
	}()

	err = s.db.
		Model(&model.Customer{}).
		Create(&customer).
		Error
	id = customer.ID
	return
}

func (s *CustomerStorage) Update(id uint, customer model.Customer) (err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error":    err,
				"id":       id,
				"customer": customer,
			}).Error("CustomerStorage.Update failed")
		}
	}()

	err = s.db.
		Model(&model.Customer{}).
		Where("id = ?", id).
		Updates(&customer).
		Error
	return
}
