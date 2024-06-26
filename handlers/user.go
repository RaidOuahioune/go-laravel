package handlers

import (
	"golang.org/x/crypto/bcrypt"

	"demo.com/hello/core/utlis"
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
	if !utlis.ValidateAndBind(ctx, &user) {
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	// Create the user in the database
	if err := db.Create(&user).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": user,
	})
}
