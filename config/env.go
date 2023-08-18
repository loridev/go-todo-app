package config

import (
	"loridev/go-todo-app/utils"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	envErr := godotenv.Load()

	if envErr != nil {
		panic(utils.DisplayError("Failed to retrieve the dotenv file variables", envErr))
	}
}
