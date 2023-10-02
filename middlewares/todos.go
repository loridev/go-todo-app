package middlewares

import (
	"loridev/go-todo-app/constants"
	"loridev/go-todo-app/types"
	"loridev/go-todo-app/utils"

	"github.com/gin-gonic/gin"
)

func IsTodoURLValid(c *gin.Context) {
	url_params := types.TodoUrlParams{}

	utils.ValidateObjectAndAbortIfErrors(c, &url_params, constants.DataFormatURLParams)
}

func CreateTodo(c *gin.Context) {
	body := types.TodoCreateRequestBody{}

	utils.ValidateObjectAndAbortIfErrors(c, &body, constants.DataFormatJSON)

	c.Set("valid_body", &body)
	c.Next()
}

func UpdateTodo(c *gin.Context) {
	body := types.TodoUpdateRequestBody{}

	utils.ValidateObjectAndAbortIfErrors(c, &body, constants.DataFormatJSON)

	c.Set("valid_body", &body)
	c.Next()
}
