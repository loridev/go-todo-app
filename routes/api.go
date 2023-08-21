package routes

import "github.com/gin-gonic/gin"

func ApiV1Router(r *gin.Engine) *gin.RouterGroup {
	router := r.Group("/api/v1")

	return router
}
