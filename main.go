package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	m "github.com/spbills/alpha-server/models"
	r "github.com/spbills/alpha-server/routers"
)

func InitDotenv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func InitRouters(router *gin.RouterGroup) {
	userRouter := router.Group("/user")
	r.SetUserRoutes(userRouter)
	applicationRouter := router.Group("/application")
	r.SetApplicationRoutes(applicationRouter)
}

func main() {
	engine := gin.Default()
	InitDotenv()
	m.ConnectDB()
	InitRouters(&engine.RouterGroup)
	engine.Run()
}
