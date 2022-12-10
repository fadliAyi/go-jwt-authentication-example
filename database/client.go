package database

import (
	"github.com/fadliAyi/go-jwt-authentication-example/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var Instance *gorm.DB
var dbError error

func Connect() {
	Instance, dbError = gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}
func Migrate() {
	Instance.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed!")
}
