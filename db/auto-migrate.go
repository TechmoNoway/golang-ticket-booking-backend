package db

import (
	"github.com/TechmoNoway/golang-ticket-booking-backend/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(&models.Event{})
}
