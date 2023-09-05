package main

import (
	"loridev/go-todo-app/config"
	"loridev/go-todo-app/models"
	"loridev/go-todo-app/routes"
	"loridev/go-todo-app/utils"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
	err := config.DB.AutoMigrate(&models.Todo{})

	if err != nil {
		panic(utils.DisplayError("Error in migration", err))
	}
}

func main() {
	defer config.CloseDB()

	r := gin.Default()

	v1Router := routes.APIV1Router(r)
	routes.TodosRouter(v1Router)

	config.RunServer(r)
}
