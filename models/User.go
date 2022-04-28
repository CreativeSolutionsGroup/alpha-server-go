package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id      int            `json:"id" gorm:"primaryKey"`
	Email   string         `json:"email"`
	Allowed pq.StringArray `gorm:"type:text[]" json:"allowed"`
}
