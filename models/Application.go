package models

import (
	"time"

	"gorm.io/gorm"
)

type Application struct {
	gorm.Model
	ID        int       `json:"ID"`
	UUID      string    `json:"UUID"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	DeletedAt time.Time `json:"DeletedAt"`
}
