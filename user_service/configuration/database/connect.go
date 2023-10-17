package connection

import (
	"fmt"
	"log"
	"os"
	configuration "user_service/configuration/enviroment"
	"user_service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func ConnectToDatabase(enviroment *configuration.Enviroment) *gorm.DB {

	// Create connection string
	uri := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", enviroment.DB_HOST, enviroment.DB_PORT, enviroment.DB_USER, enviroment.DB_PASS, enviroment.DB_NAME)

	// Generating client for communication with Postgres
	dbClient, err := gorm.Open(
		postgres.Open(uri),
		&gorm.Config{
			DryRun:         false,
			NamingStrategy: schema.NamingStrategy{SingularTable: false},
			Logger:         logger.Default.LogMode(logger.Error),
			TranslateError: true,
		},
	)
	if err != nil {
		log.Println("CONNECTION PROBLEM", err.Error())
		os.Exit(2)
	}

	// Create tables on state of structure
	if err = dbClient.AutoMigrate(
		&models.Bike{},
		&models.Brand{},
		&models.Model{},
	); err != nil {
		panic(err.Error())
	}

	return dbClient
}
