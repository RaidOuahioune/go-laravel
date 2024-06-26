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
		// "this is how to register the auth midlware"

		// refresh token belongs to the auth group

	}

	authRouter := r.Group("/auth")
	{
		authRouter.POST("/signup", handlers.SignUp)
		authRouter.POST("/login", auth.AuthMiddleware().LoginHandler)
		authRouter.GET("/refresh_token", auth.AuthMiddleware().MiddlewareFunc(), auth.AuthMiddleware().RefreshHandler)

	}
}
