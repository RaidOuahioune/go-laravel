package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

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
