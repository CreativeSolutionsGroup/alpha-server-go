package routers

import (
	"github.com/gin-gonic/gin"
	c "github.com/spbills/alpha-server/controllers"
	u "github.com/spbills/alpha-server/utils"
)

// Adds the user routes to the router.
// Users are users that are authorized to use a given set of applications.
func SetUserRoutes(router *gin.RouterGroup) {
	router.POST("", u.TokenChecker(), c.CreateUser)
	router.GET("", u.TokenChecker(), c.GetUsers)
	router.PUT("/:id", u.TokenChecker(), c.UpdateUserById)
	router.DELETE("/:id", u.TokenChecker(), c.DeleteUserById)

	router.GET("/:id", c.GetUserById)
	router.GET("/email/:email", c.GetUserByEmail)
	router.GET("/email/:email/allowed", c.GetUserApplicationsByEmail)
	router.GET("/:id/allowed", c.GetUserApplications)
}
