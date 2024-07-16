
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := SetupRouter()

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	return r
}
