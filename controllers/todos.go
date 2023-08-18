package controllers

import (
	"loridev/go-todo-app/config"
	"loridev/go-todo-app/models"
	"loridev/go-todo-app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoRequestBody struct {
	Title string `json:"title"`
}

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	config.DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	body := TodoRequestBody{}

	err := c.BindJSON(&body)

	if err != nil {
		utils.DisplayError("Error binding body", err)
	}

	todo := &models.Todo{Title: body.Title}

	dbResponse := config.DB.Create(&todo)

	if dbResponse.Error != nil {
		jsonResponse := utils.GetDBErrorJSON("insert", "todo")
		c.JSON(http.StatusInternalServerError, jsonResponse)

		return
	}

	c.JSON(http.StatusCreated, &todo)
}
