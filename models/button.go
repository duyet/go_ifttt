package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

type Event struct {
	gorm.Model
	Button string `json:"button"`
	CreatedAt time.Time
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Event{})
	return db
}