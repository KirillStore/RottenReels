package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
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
	err = db.AutoMigrate(&models.User{}, &models.Movie{}, &models.Review{}, &models.Rating{}, &models.Role{})
	if err != nil {
		return nil, fmt.Errorf("failed migration: %w", err)
	}

	var count int64
	db.Model(&models.Role{}).Count(&count)
	if count == 0 {
		roles := []models.Role{
			{Name: "employee", Desc: "default Role"},
			{Name: "admin", Desc: "admin Role"},
		}
		if err := db.Create(&roles).Error; err != nil {
			return nil, fmt.Errorf("failed creating role: %w", err)
		}
		log.Println("roles seeded")
	}
	fmt.Println("db migrated")
	DB = db
	return db, nil
}
