package main

import (
	"demo.com/hello/core/graphql"
	"demo.com/hello/core/http/auth"
	"demo.com/hello/db/migrations"
	"demo.com/hello/docs"
	"demo.com/hello/models"
	"demo.com/hello/routers"
	"demo.com/hello/services"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	services.InitSentry()
	Server()

}
func RegisterGraphQl(app *gin.Engine) {

	app.POST("/query", auth.AuthMiddleware().MiddlewareFunc(), graphql.GraphQLHandler())

}

func RegisterDocs(r *gin.Engine) {

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func Server() {

	var app = gin.Default()
	models.InitValidation()
	//migrations.PopulateDB()
	gin.SetMode(gin.ReleaseMode)

	migrations.SyncTableSchemas()

	app.Use(sentrygin.New(sentrygin.Options{}))

	auth.RegisterAuthMiddleware(app)

	routers.RegisterRoutes(app)

	RegisterGraphQl(app)

	RegisterDocs(app)

	//go job.Worker()

	app.Run() // listen and serve on 0.0.0.0:8080

}
