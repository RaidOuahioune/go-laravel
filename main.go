package main

import (
	"demo.com/hello/routers"
	"github.com/gin-gonic/gin"
)

func main() {

	Server()

}

func Server() {

	var r = gin.Default()
	routers.UserRouter(r)
	r.Run() // listen and serve on 0.0.0.0:8080

}
