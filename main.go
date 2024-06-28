package main

import (
	"demo.com/hello/core/http/auth"
	"demo.com/hello/core/job"
	"demo.com/hello/db/migrations"
	"demo.com/hello/models"
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
	models.InitValidation()

	gin.SetMode(gin.ReleaseMode)
	migrations.SyncTableSchemas()
	app.Use(sentrygin.New(sentrygin.Options{}))
	auth.RegisterAuthMiddleware(app)
	routers.UserRouter(app)

	job.Client()
	go job.Worker()

	app.Run() // listen and serve on 0.0.0.0:8080

}
