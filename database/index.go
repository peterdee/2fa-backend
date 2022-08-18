package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"backend2fa/database/models"
)

var Connection *gorm.DB

func Connect() {
	databaseName := os.Getenv("DATABASE_NAME")
	host := os.Getenv("DATABASE_HOST")
	password := os.Getenv("DATABASE_PASSWORD")
	port := os.Getenv("DATABASE_PORT")
	sslMode := os.Getenv("DATABASE_SSL_MODE")
	user := os.Getenv("DATABASE_USERNAME")
	dsn := "host=" + host + " user=" +
		user + " password=" + password + " dbname=" +
		databaseName + " port=" + port + " sslmode=" + sslMode

	instance, connectionError := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if connectionError != nil {
		log.Fatal(connectionError)
	}

	instance.AutoMigrate(&models.Users{})
	instance.AutoMigrate(&models.AuthCodes{})
	instance.AutoMigrate(&models.Passwords{})
	instance.AutoMigrate(&models.Secrets{})
	instance.AutoMigrate(&models.DeletedSecretIDs{})
	instance.AutoMigrate(&models.TokenSecrets{})

	Connection = instance
}
