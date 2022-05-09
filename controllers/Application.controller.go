package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	m "github.com/spbills/alpha-server/models"
)

func CreateApplication(c *gin.Context) {
	var applicationInput m.Application
	if err := c.ShouldBindJSON(&applicationInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if applicationInput.UUID == "" {
		applicationInput.UUID = uuid.New().String()
	}

	if err := m.Db.Create(&applicationInput).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, applicationInput)
}

func GetApplications(c *gin.Context) {
	var Applications []m.Application
	m.Db.Find(&Applications)
	c.JSON(http.StatusOK, Applications)
}

func GetApplicationById(c *gin.Context) {
	id, _ := c.Params.Get("id")

	var application m.Application
	m.Db.First(&application, id)

	c.JSON(http.StatusOK, application)
}

func UpdateApplicationById(c *gin.Context) {
	var applicationInput m.Application
	if err := c.ShouldBindJSON(&applicationInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := m.Db.Save(&applicationInput).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, applicationInput)
}

func DeleteApplicationById(c *gin.Context) {
	id, _ := c.Params.Get("id")

	m.Db.Delete(&m.Application{}, id)
}
