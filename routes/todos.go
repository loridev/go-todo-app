package routes

import (
	"loridev/go-todo-app/controllers"

	"github.com/gin-gonic/gin"
)

func TodosRouter(r *gin.RouterGroup) {
	router := r.Group("/todos")

	router.GET("", controllers.GetTodos)
	router.GET("/:id", controllers.GetTodoByID)
	router.POST("", controllers.CreateTodo)
	router.PUT("/:id", controllers.UpdateTodo)
	router.DELETE("/:id", controllers.DeleteTodo)
}
