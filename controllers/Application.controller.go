package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	m "github.com/spbills/alpha-server/models"
)

func CreateApplication(c *gin.Context) {
	var applicationInput m.Application
	if err := c.ShouldBindJSON(&applicationInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	m.Db.Save(&applicationInput)
}

func GetApplications(c *gin.Context) {
	var Applications []m.Application
	m.Db.Find(&Applications)
	c.JSON(200, Applications)
}

func GetApplicationById(c *gin.Context) {
	id, _ := c.Params.Get("id")

	var application m.Application
	m.Db.First(&application, id)

	c.JSON(200, application)
}

func UpdateApplicationById(c *gin.Context) {
	CreateApplication(c)
}

func DeleteApplicationById(c *gin.Context) {
	id, _ := c.Params.Get("id")

	m.Db.Delete(&m.Application{}, id)
}
