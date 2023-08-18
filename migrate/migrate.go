package main

import (
	"loridev/go-todo-app/config"
	"loridev/go-todo-app/models"
	"loridev/go-todo-app/utils"
)

func init() {
	config.ConnectToDB()
}

func main() {
	err := config.DB.AutoMigrate(&models.Todo{})

	if err != nil {
		panic(utils.DisplayError("Error in migration", err))
	}
}
