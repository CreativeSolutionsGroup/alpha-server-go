package routers

import (
	"github.com/gin-gonic/gin"
	c "github.com/spbills/alpha-server/controllers"
	u "github.com/spbills/alpha-server/utils"
)

func SetUserRoutes(router *gin.RouterGroup) {
	router.POST("", u.TokenChecker(), c.CreateUser)
	router.GET("", u.TokenChecker(), c.GetUsers)
	router.PUT("/:id", u.TokenChecker(), c.UpdateUserById)
	router.DELETE("/:id", u.TokenChecker(), c.DeleteUserById)
	router.GET("/:id", c.GetUserById)
	router.GET("/email/:email", c.GetUserByEmail)
	router.GET("/email/:email/allowed", c.GetUserAllowedByEmail)
	router.GET("/:id/allowed", c.GetUserAllowed)
}
