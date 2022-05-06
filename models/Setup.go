package models

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Deprecated database connection closer.
func CloseDB(db *gorm.DB) {
	// todo
}

// Connects to the database with env variables given in the README.
// Assigns the global database variable.
func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	gormDB.AutoMigrate(&User{})
	gormDB.AutoMigrate(&Application{})
	if err != nil {
		fmt.Println(err)
	}
	Db = gormDB
}

var Db *gorm.DB
