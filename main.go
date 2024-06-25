package main

import (
	"demo.com/hello/core/auth"
	"demo.com/hello/db/migrations"
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
	migrations.SyncTableSchemas()
	app.Use(sentrygin.New(sentrygin.Options{}))
	auth.RegisterAuthRoute(app)
	routers.UserRouter(app)

	app.Run() // listen and serve on 0.0.0.0:8080

}
