package storage

import (
	"ae86/internal/model"
	"ae86/pkg/logger"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type StoreStorage struct {
	db *gorm.DB
}

func NewStoreStorage(db *gorm.DB) *StoreStorage {
	return &StoreStorage{db: db}
}

func (s *StoreStorage) GetByID(id uint) (result model.Store, err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error": err,
				"id":    id,
			}).Error("StoreStorage.GetByID failed")
		}
	}()

	err = s.db.
		Model(&model.Store{}).
		Where("id = ? AND is_deleted = ?", id, false).
		First(&result).
		Error
	return
}

func (s *StoreStorage) Create(store model.Store) (id uint, err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error": err,
				"store": store,
			}).Error("StoreStorage.Create failed")
		}
	}()

	err = s.db.
		Model(&model.Store{}).
		Create(&store).
		Error
	id = store.ID
	return
}

func (s *StoreStorage) Update(id uint, store model.Store) (err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error": err,
				"id":    id,
				"store": store,
			}).Error("StoreStorage.Update failed")
		}
	}()

	err = s.db.
		Model(&model.Store{}).
		Where("id = ?", id).
		Updates(&store).
		Error
	return
}

func (s *StoreStorage) Delete(id uint) (err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error": err,
				"id":    id,
			}).Error("StoreStorage.Delete failed")
		}
	}()

	err = s.db.
		Model(&model.Store{}).
		Where("id = ?", id).
		Delete(&model.Store{}).
		Error
	return
}
