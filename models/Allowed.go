package models

import (
	"gorm.io/gorm"
)

type UserAllowed struct {
	gorm.Model
	ID            int `json:"ID"`
	UserID        int `json:"UserID"`
	User          User
	ApplicationID int `json:"ApplicationID"`
	Application   Application
}

type UserAllowedInput struct {
	ID            int `json:"ID"`
	UserID        int `json:"UserID"`
	ApplicationID int `json:"ApplicationID"`
}
