package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/otiai10/gosseract/v2"
)

func main() {
	// Define the temporary screenshot file
	tmpDir := os.TempDir()
	screenshotFile := filepath.Join(tmpDir, fmt.Sprintf("screenshot_%d.png", time.Now().Unix()))

	// Command to capture selected area (Wayland/Hyprland first, fallback to X11)
	if err := captureSelectedArea(screenshotFile); err != nil {
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

	// Copy text to clipboard (Wayland wl-copy if available, otherwise library)
	if err := copyToClipboard(text); err != nil {
		log.Fatalf("Failed to copy text to clipboard: %v", err)
	}

	fmt.Println("Text copied to clipboard successfully.")
}

func captureSelectedArea(screenshotFile string) error {
	// Prefer Wayland tools if running under Wayland/Hyprland
	if os.Getenv("WAYLAND_DISPLAY") != "" || os.Getenv("XDG_SESSION_TYPE") == "wayland" {
		// Try grim + slurp
		if has("grim") && has("slurp") {
			cmd := exec.Command("sh", "-c", fmt.Sprintf(`grim -g "$(slurp)" %q`, screenshotFile))
			return cmd.Run()
		}
		// Try grimblast (from hyprland-contrib) as a fallback
		if has("grimblast") {
			cmd := exec.Command("grimblast", "save", "area", screenshotFile)
			return cmd.Run()
		}
		// Try hyprshot (another Hyprland tool)
		if has("hyprshot") {
			cmd := exec.Command("hyprshot", "-m", "region", "-o", filepath.Dir(screenshotFile), "-f", filepath.Base(screenshotFile))
			return cmd.Run()
		}
		// If no Wayland tools, attempt scrot (may work via XWayland)
		if has("scrot") {
			cmd := exec.Command("scrot", "-s", screenshotFile)
			return cmd.Run()
		}
		return fmt.Errorf("no screenshot tool found (need grim+slurp, grimblast, hyprshot, or scrot)")
	}

	// X11 path: scrot -s
	if has("scrot") {
		cmd := exec.Command("scrot", "-s", screenshotFile)
		return cmd.Run()
	}
	return fmt.Errorf("scrot not found (install scrot or run under Wayland with grim+slurp/grimblast/hyprshot)")
}

func copyToClipboard(text string) error {
	// Wayland-first clipboard via wl-copy if present
	if (os.Getenv("WAYLAND_DISPLAY") != "" || os.Getenv("XDG_SESSION_TYPE") == "wayland") && has("wl-copy") {
		cmd := exec.Command("wl-copy")
		cmd.Stdin = strings.NewReader(text)
		if err := cmd.Run(); err == nil {
			return nil
		}
		// fall through to library on failure
	}
	// Library fallback (requires xclip or xsel on X11)
	return clipboard.WriteAll(text)
}

func has(bin string) bool {
	_, err := exec.LookPath(bin)
	return err == nil
}
