package controllers

import (
	"demo.com/hello/core/http/auth"
	"demo.com/hello/core/http/resources"
	"demo.com/hello/core/http/utlis"
	"demo.com/hello/db"
	"demo.com/hello/db/scopes"
	"demo.com/hello/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserController struct {
}

func (m *UserController) Index(ctx *gin.Context) {
	//var currentUser = auth.CurrentUser(ctx)
	var db *gorm.DB = (&db.Database{}).GetInstance()

	var pagination *utlis.Pagination = &utlis.Pagination{}

	var withCompany = ctx.Query("with_company")

	if withCompany == "true" {

		var fetchedUsers []resources.UserResource

		db.Scopes(scopes.Paginate(&fetchedUsers, pagination, ctx.Request, db)).Preload("Company").Find(&fetchedUsers)
		pagination.Rows = fetchedUsers
	} else {
		// Fetch users without Company information
		var fetchedUsers []models.User
		db.Omit("Company").Scopes(scopes.Paginate(&fetchedUsers, pagination, ctx.Request, db)).Find(&fetchedUsers)
		pagination.Rows = fetchedUsers

	}

	ctx.JSON(200, gin.H{
		"data": pagination,
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
