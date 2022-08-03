package app

import (
	"ae86/internal/model"
	"ae86/pkg/client/postgres"
	"github.com/urfave/cli/v2"
)

var migrateCmd = &cli.Command{
	Name:  "migrate",
	Usage: "migrate models to database",
	Flags: []cli.Flag{
		configFlag,
	},
	Before: loadConfig,
	Action: func(c *cli.Context) error {
		conf := getConfigFrom(c.Context)

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

		err = db.AutoMigrate(
			&model.Category{},
			&model.Product{},
			&model.Customer{},
			&model.Order{},
			&model.OrderItem{},
		)
		if err != nil {
			return err
		}

		return nil
	},
}
