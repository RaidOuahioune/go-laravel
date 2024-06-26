package routers

import (
	"demo.com/hello/core/auth"
	"demo.com/hello/handlers"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	// Simple group: v1
	userRouter := r.Group("/users")
	{
		userRouter.GET("/", auth.AuthMiddleware().MiddlewareFunc(), handlers.Index)

		userRouter.POST("/", handlers.Create)

	}
}
