package models

import (
	"gorm.io/gorm"
)

type Application struct {
	gorm.Model
	Id   int    `json:"id" gorm:"primaryKey"`
	Uuid string `json:"uuid"`
}
