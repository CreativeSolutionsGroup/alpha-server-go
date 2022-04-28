package routers

import (
	"github.com/gin-gonic/gin"
	c "github.com/spbills/alpha-server/controllers"
)

func SetUserRoutes(router *gin.RouterGroup) {
	router.POST("/user", c.CreateUser)
	router.GET("/user", c.GetUsers)
	router.GET("/user/:id", c.GetUserById)
	router.GET("/user/email/:email", c.GetUserByEmail)
	router.GET("/user/email/:email/allowed", c.GetUserAllowedByEmail)
	router.GET("/user/:id/allowed", c.GetUserAllowed)
	router.PUT("/user/:id", c.UpdateUserById)
	router.DELETE("/user/:id", c.DeleteUserById)
}
