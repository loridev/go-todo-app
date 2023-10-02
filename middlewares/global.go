package middlewares

import (
	"loridev/go-todo-app/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DoesEntityExist[T any](c *gin.Context) {
	var entity T
	id := c.Param("id")

	dbResponse := config.DB.First(&entity, id)

	if dbResponse.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Not found"})
	}

	c.Set("entity", &entity)
	c.Next()
}
