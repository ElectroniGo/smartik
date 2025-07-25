package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
	"github.com/smartik/api/internal/config"
	"github.com/smartik/api/internal/models"
	"github.com/smartik/api/internal/repository"
	"github.com/smartik/api/internal/repository/postgres"
)

func init() {
	cfg, _ := config.Load()

	if cfg.GoEnv == config.GoEnvDevelopment {
		db, err := postgres.NewConnection(cfg.PostgresURI)
		if err != nil {
			log.Fatal("Failed to connect to database:", err)
		}
		defer func() {
			sqlDb, err := db.DB()
			if err != nil {
				log.Fatal("Failed to get database instance:", err)
			}
			if err := sqlDb.Close(); err != nil {
				log.Fatal("Failed to close database connection:", err)
			}
		}()

		if err := db.AutoMigrate(models.GetAllModels()...); err != nil {
			log.Fatal("Failed to migrate database:", err)
		}

		// Seed the database with initial data if needed
		if err := repository.SeedDatabase(db); err != nil {
			log.Error("Failed to seed database:", err)
		}
	}
}

type CustomValidator struct {
	validator *validator.Validate
}

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}
