package routes

import (
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes the routes for the application.
func SetupRoutes(router *gin.Engine) {
	// Route to serve the HTML file
	router.GET("/", func(c *gin.Context) {
		c.File("./index.html") // Ensure the correct path to your HTML file
	})

	// Route to handle WAV to FLAC conversion
	router.POST("/convert", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file"})
			return
		}

		// Save the uploaded file
		err = c.SaveUploadedFile(file, "uploaded.wav")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		// Convert WAV to FLAC using ffmpeg
		cmd := exec.Command("ffmpeg", "-i", "uploaded.wav", "output.flac")
		if err := cmd.Run(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert file"})
			return
		}

		// Return the converted file
		c.File("output.flac")
	})

	// Register the WebSocket route
	router.GET("/ws", WebSocketHandler) // Registering the WebSocket handler
}
