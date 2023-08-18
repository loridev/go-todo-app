package config

import (
	"fmt"
	"loridev/go-todo-app/utils"
	"os"

	"github.com/gin-gonic/gin"
)

func RunServer(r *gin.Engine) {
	port := os.Getenv("PORT")

	serverStartErr := r.Run(fmt.Sprintf(":%s", port))

	if serverStartErr != nil {
		panic(utils.DisplayError("Failed to start the server", serverStartErr))
	}
}
