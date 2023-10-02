package utils

import (
	"encoding/json"
	"fmt"
	"loridev/go-todo-app/constants"
	"loridev/go-todo-app/types"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func DisplayError(message string, err error) string {
	return fmt.Sprintf("%s: %s", message, err)
}

func FormatValidationErrors(errors validator.ValidationErrors) map[string]string {
	result := make(map[string]string)

	for _, fieldErr := range errors {
		message := constants.TagToErrMessageMap[fieldErr.Tag()]
		lowercasedFieldName := strings.ToLower(fieldErr.Field())

		result[lowercasedFieldName] = message
	}

	return result
}

func FormatJSONTypeErrors(err *json.UnmarshalTypeError) map[string]string {
	formattedText := fmt.Sprintf("type %s cannot be used for field %s (%s)", err.Value, err.Field, err.Type.Name())

	return map[string]string{
		err.Field: formattedText,
	}
}

func validateJSON(c *gin.Context, object any) {
	if err := c.ShouldBindJSON(&object); err != nil {
		var formattedErrors map[string]string

		switch v := err.(type) {
		case validator.ValidationErrors:
			formattedErrors = FormatValidationErrors(v)
		case *json.UnmarshalTypeError:
			formattedErrors = FormatJSONTypeErrors(v)
		}

		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"errors": formattedErrors})
	}
}

func validateURL(c *gin.Context, object any) {
	url_params := types.TodoUrlParams{}

	if err := c.ShouldBindUri(&url_params); err != nil {
		formattedErrors := FormatValidationErrors(err.(validator.ValidationErrors))

		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"errors": formattedErrors})
	}
}

func ValidateObjectAndAbortIfErrors(c *gin.Context, object any, typeOfObject string) {
	switch typeOfObject {
	case constants.DataFormatJSON:
		validateJSON(c, object)
	case constants.DataFormatURLParams:
		validateURL(c, object)
	}
}
