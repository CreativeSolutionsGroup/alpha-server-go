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
	m.Db.Preload("Application").Where(&m.User{Email: email}).First(&user)

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

func GetUserApplications(c *gin.Context) {
	var userApplications []m.Application
	m.Db.Preload("Application").Find(&userApplications)

	c.JSON(http.StatusOK, userApplications)
}

func GetUserApplicationsByEmail_(email string) []m.Application {
	var user m.User
	m.Db.Preload("Applications").Find(&user, &m.User{Email: email})

	return user.Applications
}

func GetUserApplicationsByEmail(c *gin.Context) {
	email, _ := c.Params.Get("email")
	user := GetUserApplicationsByEmail_(email)

	c.JSON(http.StatusOK, user)
}
