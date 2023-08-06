package main

import (
	"fmt"
	"loridev/go-todo-app/utils"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	envErr := godotenv.Load()

	if envErr != nil {
		panic(utils.DisplayError("Failed to retrieve the dotenv file variables", envErr))
	}

	port := os.Getenv("PORT")

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	r := gin.Default()

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, dbConnErr := gorm.Open(mysql.Open(dbURL), &gorm.Config{})

	if dbConnErr != nil {
		panic(utils.DisplayError("Failed to connect to the DB", dbConnErr))
	}

	fmt.Println("Connection established with DB")

	defer func() {
		sqlDB, dbCloseErr := db.DB()

		sqlDB.Close()

		if dbCloseErr != nil {
			panic(utils.DisplayError("Failed to close the database connection", dbCloseErr))
		}

		fmt.Println("Connection closed with DB")
	}()

	serverStartErr := r.Run(fmt.Sprintf(":%s", port))

	if serverStartErr != nil {
		panic(utils.DisplayError("Failed to start the server", serverStartErr))
	}
}
