package handlers

import (
	"demo.com/hello/db"
	"demo.com/hello/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(ctx *gin.Context) {
	var db *gorm.DB = (&db.Database{}).GetInstance()
	var users []models.User
	// Find users in the database
	if err := db.Find(&users).Error; err != nil {

		ctx.JSON(500, gin.H{
			"error": "Failed to retrieve users",
		})
		return
	}
	// Return users as JSON
	ctx.JSON(200, gin.H{
		"data": users,
	})

}

func Create(ctx *gin.Context) {

	var db *gorm.DB = (&db.Database{}).GetInstance()
	var user models.User
	// Bind the request body to the user model
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	// Create the user in the database
	if err := db.Create(&user).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error": err,
		})
		return
	}
	// Return the created user as JSON
	ctx.JSON(200, gin.H{
		"data": user,
	})
}
