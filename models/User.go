package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    int    `json:"ID"`
	Email string `json:"Email"`
}
