package routers

import (
	"demo.com/hello/controllers"

	"github.com/gin-gonic/gin"
)

func TodoRouter(r *gin.Engine) {

	userRouter := r.Group("/todos")
	{

		userRouter.POST("/", (&controllers.TodoController{}).Produce)
		userRouter.GET("/", (&controllers.TodoController{}).Index)
		// "this is how to register the auth midlware"

		// refresh token belongs to the auth group

	}

}
