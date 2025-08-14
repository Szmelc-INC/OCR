#!/usr/bin/env bash
# NOT TESTED!
# setup.sh — Arch/Hyprland/Wayland OCR build helper
set -euo pipefail

need() { command -v "$1" >/dev/null 2>&1; }

# Require pacman
if ! need pacman; then
  echo "error: this script targets Arch-based systems (needs pacman)"; exit 1
fi

# Packages (install if available in repos)
pkgs=(
  go
  tesseract
  leptonica
  pkgconf
  grim
  slurp
  wl-clipboard
  hyprshot
  # hyprland-contrib   # grimblast
  scrot              # XWayland fallback
  xclip xsel         # X11 clipboard fallback
)

echo ":: installing dependencies (skipping missing pkgs)…"
sudo pacman -Sy --noconfirm
for p in "${pkgs[@]}"; do
  if pacman -Si "$p" &>/dev/null; then
    sudo pacman -S --needed --noconfirm "$p"
  else
    echo "  - $p not found in repos, skipping"
  fi
done

# Optional: English language data for Tesseract (skip if unavailable)
if pacman -Si tesseract-data-eng &>/dev/null; then
  sudo pacman -S --needed --noconfirm tesseract-data-eng
fi

# Ensure we’re in the project root (where main.go resides)
# (no-op if already correct)
if ! ls *.go >/dev/null 2>&1; then
  echo "error: run this script in the Go project directory (with main.go)"; exit 1
fi

# Initialize go.mod if missing (uses module name 'ocr' to match your setup)
if [ ! -f go.mod ]; then
  go mod init ocr
fi

# Handle flaky DNS to Go proxy: default to direct unless GOPROXY set
export GOPROXY="${GOPROXY:-direct}"
# Disable sumdb only if user didn’t set it explicitly
export GOSUMDB="${GOSUMDB:-off}"

# Resolve deps and build
go mod tidy
go build -o ocr

echo ":: build complete -> ./ocr"
echo ":: tip: On Wayland/Hyprland you'll typically use grim+slurp or grimblast; on XWayland fallback is scrot."
