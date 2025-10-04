package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router with default middleware
	r := gin.Default()

	// Health check endpoint
	r.GET("/v1/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Start server on port 8080
	r.Run(":8080")
}
