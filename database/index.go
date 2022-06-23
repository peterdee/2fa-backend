package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func Connect() {
	databaseName := os.Getenv("DATABASE_NAME")
	host := os.Getenv("DATABASE_HOST")
	password := os.Getenv("DATABASE_PASSWORD")
	port := os.Getenv("DATABASE_PORT")
	user := os.Getenv("DATABASE_USERNAME")
	dsn := "host=" + host + " user=" +
		user + " password=" + password + " dbname=" +
		databaseName + " port=" + port + " sslmode=disable"
	instance, connectionError := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	fmt.Println(dsn, instance)
	if connectionError != nil {
		log.Fatal(connectionError)
	}

	Connection = instance
}