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
	r.SetUserRoutes(router)
}

func main() {
	engine := gin.Default()
	InitRouters(&engine.RouterGroup)
	InitDotenv()
	m.ConnectDB()
	engine.Run()
}
