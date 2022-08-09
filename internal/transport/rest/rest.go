package rest

import (
	"ae86/internal/container"
	"ae86/pkg/logger"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func Start(config Config, container *container.RestContainer) {
	app := fiber.New(fiber.Config{
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		DisableStartupMessage: true,
	})

	// set middlewares
	RegisterRoutes(app, container)

	address := config.Address()

	var err error
	if config.TLSEnable {
		err = app.ListenTLS(address, config.CertFile, config.KeyFile)
	} else {
		err = app.Listen(address)
	}
	if err != nil {
		logger.Log.Fatalf("failed to start rest server: %v", err)
	}
}
