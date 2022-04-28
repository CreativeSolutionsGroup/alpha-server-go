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

}

func GetUserById(c *gin.Context) {
	id, _ := c.Params.Get("id")

	var user m.User
	m.Db.Find(&user, id)

	c.JSON(200, user)
}

func UpdateUserById(c *gin.Context) {

}

func DeleteUserById(c *gin.Context) {

}
