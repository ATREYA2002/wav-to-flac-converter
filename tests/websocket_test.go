package tests

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestWebSocketConversion(t *testing.T) {
	// Load WAV data from a file
	wavData, err := ioutil.ReadFile("BAK.wav") // Ensure the correct path to your WAV file
	if err != nil {
		t.Fatalf("Failed to read WAV file: %v", err)
	}

	url := "ws://localhost:8080/ws" // Ensure this matches your WebSocket endpoint

	// Attempt to establish a WebSocket connection with a timeout
	var ws *websocket.Conn
	for i := 0; i < 5; i++ { // Try up to 5 times
		ws, _, err = websocket.DefaultDialer.Dial(url, nil)
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second) // Wait for 1 second before retrying
	}

	if err != nil {
		t.Fatalf("WebSocket connection failed: %v", err)
	}
	defer ws.Close()

	// Simulate sending WAV data
	err = ws.WriteMessage(websocket.BinaryMessage, wavData)
	if err != nil {
		t.Fatalf("Failed to send WAV data: %v", err)
	}

	// Read the FLAC response
	_, flacData, err := ws.ReadMessage()
	if err != nil {
		t.Fatalf("Failed to read FLAC response: %v", err)
	}

	// Validate FLAC data
	if len(flacData) == 0 {
		t.Fatal("Expected non-empty FLAC data")
	}

	// Additional checks (e.g., validate FLAC format if necessary)
}
