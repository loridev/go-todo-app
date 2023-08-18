package main

import (
	"loridev/go-todo-app/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnvVariables()
	config.ConnectToDB()
	defer config.CloseDB()

	r := gin.Default()
	config.RunServer(r)
}
