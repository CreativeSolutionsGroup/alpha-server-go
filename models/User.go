package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           int           `json:"ID"`
	Email        string        `json:"Email"`
	Password     string        `json:"Password"`
	CreatedAt    time.Time     `json:"CreatedAt"`
	UpdatedAt    time.Time     `json:"UpdatedAt"`
	DeletedAt    time.Time     `json:"DeletedAt"`
	Applications []UserAllowed `gorm:"foreignkey:UserID" json:"Applications"`
}
