package service

import (
	"ae86/internal/model"
	"ae86/internal/service/adapter"
	"ae86/pkg/logger"
	"github.com/sirupsen/logrus"
	"log"
)

type ManagerService struct {
	storage adapter.StorageContainer
}

func NewManagerService(storage adapter.StorageContainer) *ManagerService {
	return &ManagerService{storage: storage}
}

func (s *ManagerService) GetByID(id uint) (result model.Manager, err error) {
	defer func() {
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error": err,
				"id":    id,
			}).Error("ManagerService.GetByID failed")
		}
	}()

	return s.storage.Manager().GetByID(id)
}

func (s *ManagerService) GetManager() (result model.Manager) {
	result, err := s.storage.Manager().GetByID(0)
	if err != nil {
		log.Println(err)
	}
	return result
}
