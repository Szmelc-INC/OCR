package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/atotto/clipboard"
	"github.com/otiai10/gosseract/v2"
)

func main() {
	// Define the temporary screenshot file
	tmpDir := os.TempDir()
	screenshotFile := filepath.Join(tmpDir, fmt.Sprintf("screenshot_%d.png", time.Now().Unix()))

	// Command to capture selected area
	cmd := exec.Command("scrot", "-s", screenshotFile)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to take screenshot: %v", err)
	}

	// Check if the screenshot file exists
	if _, err := os.Stat(screenshotFile); os.IsNotExist(err) {
		log.Fatalf("Screenshot file does not exist: %v", err)
	}

	// Perform OCR
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage(screenshotFile)

	text, err := client.Text()
	if err != nil {
		log.Fatalf("Failed to perform OCR: %v", err)
	}

	// Copy text to clipboard
	err = clipboard.WriteAll(text)
	if err != nil {
		log.Fatalf("Failed to copy text to clipboard: %v", err)
	}

	fmt.Println("Text copied to clipboard successfully.")
}
