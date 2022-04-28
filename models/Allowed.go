package models

import (
	"gorm.io/gorm"
)

type UserAllowed struct {
	gorm.Model
	Id            int `json:"id" gorm:"primaryKey"`
	UserId        int `json:"user_id"`
	User          User
	ApplicationId int `json:"application_id"`
	Application   Application
}

type UserAllowedInput struct {
	Id            int `json:"id" gorm:"primaryKey"`
	UserId        int `json:"user_id"`
	ApplicationId int `json:"application_id"`
}
