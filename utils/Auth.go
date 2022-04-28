package utils

import (
	"os"

	"github.com/gin-gonic/gin"
)

func TokenChecker() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.Request.Header.Get("Authorization")
		sysKey := os.Getenv("API_KEY")
		if apiKey == "" || apiKey != sysKey {
			c.AbortWithStatus(401)
			return
		}
		c.Next()
	}
}
