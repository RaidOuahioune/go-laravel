package auth

import (
	"log"

	"demo.com/hello/core/auth/forms"
	"demo.com/hello/core/utlis"
	"demo.com/hello/db"
	"demo.com/hello/models"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	IdentityKey = "ID"
)

func AuthMiddleware() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(InitJwtParams())
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())

	}
	return authMiddleware
}
func RegisterAuthMiddleware(r *gin.Engine) {
	authMiddleware := AuthMiddleware()
	r.Use(handlerMiddleWare(authMiddleware))

}

func handlerMiddleWare(authMiddleware *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return func(context *gin.Context) {
		errInit := authMiddleware.MiddlewareInit()
		if errInit != nil {
			log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
		}
	}
}

func payloadFunc() func(data interface{}) jwt.MapClaims {
	return func(data interface{}) jwt.MapClaims {
		if v, ok := data.(*models.User); ok {
			return jwt.MapClaims{
				IdentityKey: v.ID,
			}
		}
		return jwt.MapClaims{}
	}
}

func identityHandler() func(c *gin.Context) interface{} {
	return func(c *gin.Context) interface{} {
		claims := jwt.ExtractClaims(c)
		var db *gorm.DB = (&db.Database{}).GetInstance()
		var user models.User

		var err = db.First(&user, claims[IdentityKey])
		if err.Error != nil {
			c.JSON(401, gin.H{
				"error": "User not found from identyHandler",
			})
		}
		return &user
	}
}

func authenticator() func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		var loginVals forms.LoginForm
		utlis.ValidateAndBind(c, &loginVals)
		var db *gorm.DB = (&db.Database{}).GetInstance()
		var user models.User

		var err = db.Find(&user, "email", loginVals.Email)

		if err.Error != nil {
			return nil, jwt.ErrFailedAuthentication
		}
		var passError = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginVals.Password))
		if passError != nil {
			return nil, jwt.ErrFailedAuthentication
		}
		return &user, nil

	}
}

func authorizator() func(data interface{}, c *gin.Context) bool {
	return func(data interface{}, c *gin.Context) bool {
		return true
	}
}

func unauthorized() func(c *gin.Context, code int, message string) {
	return func(c *gin.Context, code int, message string) {
		c.JSON(code, gin.H{
			"code":    code,
			"message": message,
		})
	}
}
