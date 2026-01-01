package migration

import (
	"carta-ecommerce-api/internal/config"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"gorm.io/gorm"
)

func Run(db *gorm.DB, cfg config.DatabaseConfig) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		cfg.Name,
		driver,
	)
	if err != nil {
		return err
	}

	return m.Up()
}
