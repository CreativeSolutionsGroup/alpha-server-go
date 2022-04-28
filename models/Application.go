package models

import (
	"gorm.io/gorm"
)

type Application struct {
	gorm.Model
	ID   int    `json:"ID"`
	UUID string `json:"UUID"`
}
