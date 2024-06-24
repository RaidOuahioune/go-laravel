package main

import (
	"demo.com/hello/routers"
	"demo.com/hello/services"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func main() {

	services.InitSentry()
	Server()

}

func Server() {

	var app = gin.Default()
	gin.SetMode(gin.ReleaseMode)
	app.Use(sentrygin.New(sentrygin.Options{}))
	routers.UserRouter(app)
	app.Run() // listen and serve on 0.0.0.0:8080

}
