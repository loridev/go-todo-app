package config

import (
	"fmt"
	"loridev/go-todo-app/utils"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	LoadEnvVariables()

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	var dbConnErr error

	dbURL := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		dbHost,
		dbUsername,
		dbPassword,
		dbName,
		dbPort,
	)

	DB, dbConnErr = gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if dbConnErr != nil {
		CloseDB()

		panic(utils.DisplayError("Failed to connect to the DB", dbConnErr))
	}

	fmt.Println("Connection established with DB")
}

func CloseDB() {
	sqlDB, dbCloseErr := DB.DB()

	sqlDB.Close()

	if dbCloseErr != nil {
		panic(utils.DisplayError("Failed to close the database connection", dbCloseErr))
	}

	fmt.Println("Connection closed with DB")
}
