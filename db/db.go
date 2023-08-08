package db

import (
	"fmt"

	"github.com/viniciusarambul/transaction/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func SetupDB() (*gorm.DB, error) {
	cfg := config.LoadEnvVars()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DatabaseHost, cfg.DatabaseUsername, cfg.DatabasePassword, cfg.DatabaseName, cfg.DatabasePort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger.Discard,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   fmt.Sprintf("%s.", cfg.DatabaseSchema),
			SingularTable: false,
		},
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}
