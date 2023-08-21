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

	v1Router := routes.ApiV1Router(r)
	routes.TodosRouter(v1Router)

	config.RunServer(r)
}
