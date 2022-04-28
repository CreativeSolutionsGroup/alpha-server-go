package models

import (
	"time"

	"gorm.io/gorm"
)

type UserAllowed struct {
	gorm.Model
	ID            int `json:"ID"`
	UserID        int `json:"UserID"`
	User          User
	ApplicationID int `json:"ApplicationID"`
	Application   Application
	CreatedAt     time.Time `json:"CreatedAt"`
	UpdatedAt     time.Time `json:"UpdatedAt"`
	DeletedAt     time.Time `json:"DeletedAt"`
}

type UserAllowedInput struct {
	ID            int `json:"ID"`
	UserID        int `json:"UserID"`
	ApplicationID int `json:"ApplicationID"`
}
