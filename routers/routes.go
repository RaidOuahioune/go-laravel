package routers

import "github.com/gin-gonic/gin"

var routes = []func(r *gin.Engine){

	UserRouter,
	TodoRouter,
}

func RegisterRoutes(r *gin.Engine) {

	for _, route := range routes {
		route(r)
	}
}
