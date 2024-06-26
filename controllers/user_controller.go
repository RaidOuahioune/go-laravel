package controllers

import (
	"demo.com/hello/core/auth"
	"demo.com/hello/core/utlis"
	"demo.com/hello/db"
	"demo.com/hello/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserController struct {
}

func (m *UserController) Index(ctx *gin.Context) {
	var currentUser = auth.CurrentUser(ctx)
	var db *gorm.DB = (&db.Database{}).GetInstance()
	var users []models.User
	// ctx.JSON(200, gin.H{
	// 	"sql": db.ToSQL(func(tx *gorm.DB) *gorm.DB {
	// 		return tx.Find(&users)

	// 	})})
	// Find users in the database
	db.Find(&users)
	// Return users as JSON
	ctx.JSON(200, gin.H{
		"data":         users,
		"current_user": currentUser,
	})

}

func (m *UserController) SignUp(ctx *gin.Context) {

	var db *gorm.DB = (&db.Database{}).GetInstance()
	var user models.User
	// Bind the request body to the user model
	if !utlis.ValidateAndBind(ctx, &user) {
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	db.First(&user, "email", user.Email)
	if user.ID != 0 {
		ctx.JSON(400, gin.H{
			"error": "Email already exists",
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
	var token, expires, err = auth.AuthMiddleware().TokenGenerator(&user)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data":   user,
		"token":  token,
		"expire": expires,
	})
}
