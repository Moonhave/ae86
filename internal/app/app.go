package app

import (
	"ae86/config"
	"ae86/internal/container"
	"ae86/internal/transport"
	"ae86/internal/transport/rest"
	"ae86/pkg/client/postgres"
	"ae86/pkg/logger"
)

func Run(conf config.Config) error {
	db, err := postgres.Connect(postgres.Config{
		Username: conf.DB.Username,
		Password: conf.DB.Password,
		Host:     conf.DB.Host,
		Port:     conf.DB.Port,
		Database: conf.DB.Database,
		SSLMode:  conf.DB.SSLMode,
	})
	if err != nil {
		return err
	}

	logger.Log.Info("connected to database...")

	storage := container.NewStorageContainer(db)
	service := container.NewServiceContainer(storage)
	controller := container.NewRestContainer(service)

	transportConfig := rest.Config{
		Host:      conf.HTTP.Host,
		Port:      conf.HTTP.Port,
		TLSEnable: conf.HTTP.TLSEnable,
		CertFile:  conf.HTTP.CertFile,
		KeyFile:   conf.HTTP.KeyFile,
		BotToken:  conf.Bot.Token,
	}

	return transport.Start(transportConfig, controller)
}
