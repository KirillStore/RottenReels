package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"test_go/config"
	"test_go/models"
)

var DB *gorm.DB

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.Username, cfg.Database.Dbname, cfg.Database.Password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}
	err = db.AutoMigrate(&models.User{}, &models.Movie{}, models.Review{}, models.Rating{})
	if err != nil {
		return nil, fmt.Errorf("failed migration: %w", err)
	}
	fmt.Println("db migrated")
	DB = db
	return db, nil
}
