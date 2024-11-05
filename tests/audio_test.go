package tests

import (
	"os"
	"testing"
	"wavtoflacconverter/services"
)

func TestConvertWAVToFLAC(t *testing.T) {
	// Read WAV file using os
	wavData, err := os.ReadFile("BAK.wav")
	if err != nil {
		t.Fatalf("Failed to read WAV file: %v", err)
	}

	// Call the conversion function
	flacData, err := services.ConvertWAVToFLAC(wavData)
	if err != nil {
		t.Fatalf("Conversion failed: %v", err)
	}

	// Validate FLAC data
	if len(flacData) == 0 {
		t.Fatal("Expected non-empty FLAC data")
	}
}
