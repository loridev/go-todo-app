package utils

import (
	"fmt"
	"loridev/go-todo-app/types"
	"net/http"
)

func GetDBErrorJSON(operation DBOperation, entityName string) *types.RestError {
	err := ValidateOperation(operation)

	if err != nil {
		panic(err)
	}

	message := fmt.Sprintf("Error while performing the operation \"%s\" on %s", operation, entityName)

	return &types.RestError{
		Error: message,
		Status: http.StatusInternalServerError,
	}
}

func BadRequestError(message string) *types.RestError {
	return &types.RestError{
		Error: message,
		Status: http.StatusBadRequest,
	}
}
