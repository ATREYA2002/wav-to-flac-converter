// services/audio.go
package services

import (
	"bytes"
	"os/exec"
)

func ConvertWAVToFLAC(wavData []byte) ([]byte, error) {
	// Set up the command to run ffmpeg
	cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-f", "flac", "pipe:1")
	cmd.Stdin = bytes.NewReader(wavData)

	var outBuf bytes.Buffer
	cmd.Stdout = &outBuf

	// Run the command and capture any errors
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return outBuf.Bytes(), nil
}
