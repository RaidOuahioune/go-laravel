package utlis

import (
	"fmt"

	"demo.com/hello/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func FormatValidationErrors(err error) map[string][]string {
	errs := make(map[string][]string)
	for _, err := range err.(validator.ValidationErrors) {
		field := err.Field()
		errMsg := fmt.Sprintf("failed on the '%s' tag", err.Tag())
		errs[field] = append(errs[field], errMsg)
	}
	return errs
}

func ValidateAndBind(ctx *gin.Context, obj interface{}) bool {

	var err = ctx.BindJSON(obj)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return false
	}

	err = models.Validate.Struct(obj)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": FormatValidationErrors(err),
		})
		return false
	}
	return true
}

// ctx.JSON(200, gin.H{
// 	"sql": db.ToSQL(func(tx *gorm.DB) *gorm.DB {
// 		return tx.Find(&users)

// 	})})
// Find users in the database
