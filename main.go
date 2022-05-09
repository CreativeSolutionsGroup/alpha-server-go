package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	m "github.com/spbills/alpha-server/models"
	r "github.com/spbills/alpha-server/routers"
)

func InitDotenv() {
	if os.Getenv("GIN_MODE") != "release" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func InitRouters(router *gin.RouterGroup) {
	userRouter := router.Group("/user")
	r.SetUserRoutes(userRouter)
	applicationRouter := router.Group("/application")
	r.SetApplicationRoutes(applicationRouter)
}

func InitDatabases() {
	err := m.ConnectDB()

	if err != nil {
		fmt.Println("Failed to connect to database. Trying again.")
		time.Sleep(time.Second)
		InitDatabases()
	}
}

func main() {
	engine := gin.Default()
	InitDotenv()
	InitDatabases()
	InitRouters(&engine.RouterGroup)

	fmt.Println("Successfully connected to database and initialized routers.")
	engine.Run()
}
