package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root@tcp(database:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err == nil {
		fmt.Println("raid no error")
	} else {
		fmt.Println(err.Error())

	}

	db.Exec("create table lol(id integer primary key);")
	var r = gin.Default()
	r.GET("/ping", func(c *gin.Context) {

		var response string = ""
		for {
			response += "Raid is learning loops this in an update to see if comppose it aware of it."
			break
		}
		c.JSON(200, gin.H{
			"message": response,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
