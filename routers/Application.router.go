package routers

import (
	"github.com/gin-gonic/gin"
	c "github.com/spbills/alpha-server/controllers"
	u "github.com/spbills/alpha-server/utils"
)

func SetApplicationRoutes(router *gin.RouterGroup) {
	router.POST("/application", u.TokenChecker(), c.CreateApplication)
	router.PUT("/application/:id", u.TokenChecker(), c.UpdateApplicationById)
	router.DELETE("/application/:id", u.TokenChecker(), c.DeleteApplicationById)

	router.GET("/application", c.GetApplications)
	router.GET("/application/:id", c.GetApplicationById)
}
