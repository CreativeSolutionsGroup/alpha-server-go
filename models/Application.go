package models

import (
	"time"

	"gorm.io/gorm"
)

// An application is a certain app that a user is allowed to access.
// Many 2 many with users.
type Application struct {
	gorm.Model
	ID        int        `json:"ID"`
	UUID      string     `json:"UUID"`
	Users     []User     `gorm:"many2many:user_alloweds;" json:"Users"`
	CreatedAt *time.Time `json:"CreatedAt"`
	UpdatedAt *time.Time `json:"UpdatedAt"`
	DeletedAt *time.Time `json:"DeletedAt"`
}
