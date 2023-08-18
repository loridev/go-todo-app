package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetDbErrorJson(operation DBOperation, entityName string) gin.H {
	err := ValidateOperation(operation)

	if err != nil {
		panic(err)
	}

	message := fmt.Sprintf("Error while performing the operation \"%s\" on %s", operation, entityName)

	return gin.H{"Error": message}
}
