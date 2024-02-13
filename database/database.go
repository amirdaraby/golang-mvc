package database

import (
	"fmt"

	"github.com/amirdaraby/golang-mvc/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbConnection *gorm.DB

func Init(cfg *config.Database) error {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran", cfg.Host, cfg.Username, cfg.Password, cfg.Name, cfg.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	sqlDB, err := db.DB()

	if err != nil {
		return err
	}

	err = sqlDB.Ping()

	if err != nil {
		return err
	}

	DbConnection = db

	return nil
}
