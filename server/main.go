
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Setup router
	r := SetupRouter()

	// Start server
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

// SetupRouter configures the Gin router with routes and middleware
func SetupRouter() *gin.Engine {
	// Initialize Gin router
	r := gin.Default()

	// Example route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	return r
}
