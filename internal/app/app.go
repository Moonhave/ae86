package app

import (
	"ae86/config"
	"ae86/internal/container"
	"ae86/internal/transport"
	"ae86/internal/transport/bot"
	"ae86/internal/transport/rest"
	"ae86/pkg/client/postgres"
	"ae86/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func Run(conf config.Config) {
	db, err := postgres.Connect(postgres.Config{
		Username: conf.DB.Username,
		Password: conf.DB.Password,
		Host:     conf.DB.Host,
		Port:     conf.DB.Port,
		Database: conf.DB.Database,
		SSLMode:  conf.DB.SSLMode,
	})
	if err != nil {
		logger.Log.Fatalf("failed to connect to database: %v", err)
	}

	logger.Log.Info("connected to database...")

	storage := container.NewStorageContainer(db)
	service := container.NewServiceContainer(storage)
	restControllers := container.NewRestContainer(service)
	botHandlers := container.NewHandlerContainer(service)

	transportConfig := transport.Config{
		Rest: rest.Config{
			Host:      conf.HTTP.Host,
			Port:      conf.HTTP.Port,
			TLSEnable: conf.HTTP.TLSEnable,
			CertFile:  conf.HTTP.CertFile,
			KeyFile:   conf.HTTP.KeyFile,
		},
		Bot: bot.Config{
			Token:             conf.Bot.Token,
			LongPollerTimeout: conf.Bot.LongPollerTimeout,
		},
	}

	transport.Start(transportConfig, restControllers, botHandlers)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	// graceful shutdown...
}
