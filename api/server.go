package api

import (
	"github.com/gin-gonic/gin"
)

type HelloRequest struct {
	Name string `json:"name" binding:"required"`
}

func RunHTTPServer() {
	r := gin.Default()

	r.POST("/hello", func(c *gin.Context) {
		var request HelloRequest
		if err := c.BindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": "Name parameter is required"})
			return
		}
		c.JSON(200, gin.H{
			"message": "Hello " + request.Name,
		})
	})

	r.Run(":8080")
}
