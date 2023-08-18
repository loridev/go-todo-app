package routes

import (
	"loridev/go-todo-app/controllers"

	"github.com/gin-gonic/gin"
)

func TodosRouter(r *gin.Engine) {
	routes := r.Group("/api/v1/todos")

	routes.GET("", controllers.GetTodos)
	routes.POST("", controllers.CreateTodo)
}
