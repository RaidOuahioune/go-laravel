package routers

import (
	"demo.com/hello/handlers"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	// Simple group: v1
	userRouter := r.Group("/users")
	{
		userRouter.GET("/", handlers.Index)

		userRouter.POST("/", handlers.Create)

	}
}
