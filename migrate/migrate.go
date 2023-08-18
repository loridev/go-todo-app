package main

import (
	"loridev/go-todo-app/config"
	"loridev/go-todo-app/models"
)

func init() {
	config.ConnectToDB()
}

func main() {
	config.DB.AutoMigrate(&models.Todo{})
}
