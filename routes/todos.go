package routes

import (
	"loridev/go-todo-app/controllers"
	"loridev/go-todo-app/middlewares"
	"loridev/go-todo-app/models"

	"github.com/gin-gonic/gin"
)

func TodosRouter(r *gin.RouterGroup) {
	router := r.Group("/todos")
	router.GET("", controllers.GetTodos)
	router.GET(
		"/:id",
		middlewares.IsTodoURLValid,
		middlewares.DoesEntityExist[models.Todo],
		controllers.GetTodoByID,
	)
	router.POST("", middlewares.CreateTodo, controllers.CreateTodo)
	router.PUT(
		"/:id",
		middlewares.IsTodoURLValid,
		middlewares.DoesEntityExist[models.Todo],
		middlewares.UpdateTodo,
		controllers.UpdateTodo,
	)
	router.DELETE(
		"/:id",
		middlewares.IsTodoURLValid,
		middlewares.DoesEntityExist[models.Todo],
		controllers.DeleteTodo,
	)
}
