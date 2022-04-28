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
	c.JSON(200, users)
}

func GetUserById(c *gin.Context) {
	id, _ := c.Params.Get("id")

	var user m.User
	m.Db.First(&user, id)

	c.JSON(200, user)
}

func UpdateUserById(c *gin.Context) {
	CreateUser(c)
}

func DeleteUserById(c *gin.Context) {
	id, _ := c.Params.Get("id")

	m.Db.Delete(&m.User{}, id)
}

func GetUserAllowed(c *gin.Context) {
	// Join users with applications using allowed table
	var userAllowed []m.UserAllowed
	m.Db.Preload("User").Preload("Application").Find(&userAllowed)

	c.JSON(200, userAllowed)
}
