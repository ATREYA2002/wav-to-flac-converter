package routes

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes the routes for the application.
func SetupRoutes(router *gin.Engine) {
	// Route to serve the HTML file
	router.GET("/", func(c *gin.Context) {
		c.File("./index.html") // Ensure the correct path to your HTML file
	})

	router.GET("/convert", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the WAV to FLAC conversion service!"})
	})

	// Register the WebSocket route
	router.GET("/ws", WebSocketHandler) // Registering the WebSocket handler
}
