package main

import (
	"loridev/go-todo-app/config"
	"loridev/go-todo-app/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}

func main() {
	defer config.CloseDB()

	r := gin.Default()

	routes.TodosRouter(r)

	config.RunServer(r)
}
