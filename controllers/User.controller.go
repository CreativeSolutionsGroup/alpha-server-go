package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	m "github.com/spbills/alpha-server/models"
)

func CreateUser(c *gin.Context) {
	var userInput m.User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	m.Db.Save(&userInput)
}

func GetUsers(c *gin.Context) {
	var users []m.User
	m.Db.Find(&users)
	c.JSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {
	id, _ := c.Params.Get("id")

	var user m.User
	m.Db.First(&user, id)

	c.JSON(http.StatusOK, user)
}

func GetUserByEmail_(email string) m.User {
	var user m.User
	m.Db.Preload("UserAllowed").Where("email = ?", email).First(&user)

	return user
}

func GetUserByEmail(c *gin.Context) {
	email, _ := c.Params.Get("email")
	user := GetUserByEmail_(email)

	c.JSON(http.StatusOK, user)
}

func UpdateUserById(c *gin.Context) {
	CreateUser(c)
}

func DeleteUserById(c *gin.Context) {
	id, _ := c.Params.Get("id")

	m.Db.Delete(&m.User{}, id)
}

func GetUserAllowed(c *gin.Context) {
	var userAllowed []m.UserAllowed
	m.Db.Preload("Application").Find(&userAllowed)

	c.JSON(http.StatusOK, userAllowed)
}

func GetUserAllowedByEmail_(email string) []m.UserAllowed {
	var userAllowed []m.UserAllowed
	m.Db.Preload("Application").Joins("User", m.Db.Where(&m.User{Email: email})).Find(&userAllowed)

	return userAllowed
}

func GetUserAllowedByEmail(c *gin.Context) {
	email, _ := c.Params.Get("email")
	user := GetUserAllowedByEmail_(email)

	c.JSON(http.StatusOK, user)
}
