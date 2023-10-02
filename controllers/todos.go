package controllers

import (
	"loridev/go-todo-app/config"
	"loridev/go-todo-app/models"
	"loridev/go-todo-app/types"
	"loridev/go-todo-app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []models.Todo

	config.DB.Find(&todos)

	c.JSON(http.StatusOK, todos)
}

func GetTodoByID(c *gin.Context) {
	todo := c.MustGet("entity").(*models.Todo)

	c.JSON(http.StatusOK, todo)
}

func CreateTodo(c *gin.Context) {
	body := c.MustGet("valid_body").(*types.TodoCreateRequestBody)

	todo := &models.Todo{Title: body.Title}

	dbResponse := config.DB.Create(&todo)

	if dbResponse.Error != nil {
		jsonResponse := utils.GetDBErrorJSON("insert", "todo")

		c.AbortWithStatusJSON(http.StatusInternalServerError, jsonResponse)
	}

	c.JSON(http.StatusCreated, &todo)
}

func UpdateTodo(c *gin.Context) {
	todo := c.MustGet("entity").(*models.Todo)

	body := c.MustGet("valid_body").(*types.TodoUpdateRequestBody)

	dbResponse := config.DB.Model(&todo).Updates(body)

	if dbResponse.Error != nil {
		jsonResponse := utils.GetDBErrorJSON("update", "todo")
		c.JSON(http.StatusInternalServerError, jsonResponse)

		return
	}

	c.JSON(http.StatusOK, &todo)
}

func DeleteTodo(c *gin.Context) {
	todo := c.MustGet("entity").(*models.Todo)

	dbResponse := config.DB.Delete(&todo)

	if dbResponse.Error != nil {
		jsonResponse := utils.GetDBErrorJSON("delete", "todo")
		c.JSON(http.StatusInternalServerError, jsonResponse)

		return
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}
