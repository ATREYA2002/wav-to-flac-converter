package routes

import (
	"log"
	"net/http"
	"wavtoflacconverter/services" // Import the services package

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func WebSocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Failed to set up WebSocket: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to set up WebSocket"})
		return
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error closing WebSocket connection: %v", err)
		}
	}()

	for {
		// Read the incoming WAV audio data
		_, audioData, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break // Exit loop on read error
		}

		// Validate the incoming audio data
		if len(audioData) == 0 {
			log.Println("Received empty audio data")
			conn.WriteMessage(websocket.TextMessage, []byte("Received empty audio data"))
			continue // Continue listening for more messages
		}

		// Convert WAV to FLAC
		flacData, err := services.ConvertWAVToFLAC(audioData)
		if err != nil {
			log.Printf("Error converting audio: %v", err)
			errMsg := "Error converting audio: " + err.Error()
			conn.WriteMessage(websocket.TextMessage, []byte(errMsg))
			continue // Continue listening for more messages
		}

		// Send a success message before sending the converted data
		successMessage := "Conversion successful! Sending FLAC data."
		err = conn.WriteMessage(websocket.TextMessage, []byte(successMessage))
		if err != nil {
			log.Printf("Error sending success message: %v", err)
			break // Exit loop on write error
		}

		// Send the converted FLAC data back to the client
		err = conn.WriteMessage(websocket.BinaryMessage, flacData)
		if err != nil {
			log.Printf("Error sending FLAC data: %v", err)
			break // Exit loop on write error
		}
	}
}
