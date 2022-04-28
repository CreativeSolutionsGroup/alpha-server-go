package routers

import (
	"github.com/gin-gonic/gin"
	c "github.com/spbills/alpha-server/controllers"
	u "github.com/spbills/alpha-server/utils"
)

func SetApplicationRoutes(router *gin.RouterGroup) {
	router.POST("", u.TokenChecker(), c.CreateApplication)
	router.PUT("/:id", u.TokenChecker(), c.UpdateApplicationById)
	router.DELETE("/:id", u.TokenChecker(), c.DeleteApplicationById)

	router.GET("", c.GetApplications)
	router.GET("/:id", c.GetApplicationById)
}
