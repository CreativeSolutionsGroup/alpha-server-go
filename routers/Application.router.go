package routers

import (
	"github.com/gin-gonic/gin"
	c "github.com/spbills/alpha-server/controllers"
)

func SetApplicationRoutes(router *gin.RouterGroup) {
	router.POST("/application", c.CreateApplication)
	router.GET("/application", c.GetApplications)
	router.GET("/application/:id", c.GetApplicationById)
	router.PUT("/application/:id", c.UpdateApplicationById)
	router.DELETE("/application/:id", c.DeleteApplicationById)
}
