package models

import (
	"time"

	"gorm.io/gorm"
)

// Users are users that are authorized to use a given set of applications.
type User struct {
	gorm.Model
	ID           int           `json:"ID"`
	Email        string        `json:"Email"`
	CreatedAt    *time.Time    `json:"CreatedAt"`
	UpdatedAt    *time.Time    `json:"UpdatedAt"`
	DeletedAt    *time.Time    `json:"DeletedAt"`
	Applications []Application `gorm:"many2many:user_alloweds;" json:"Applications"`
}
