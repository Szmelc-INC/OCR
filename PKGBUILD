# Maintainer: SilverX <serainox@gmail.com>
pkgname=ocr-git
_repo=OCR
pkgver=r0.0000000
pkgrel=1
pkgdesc="Screenshot area + OCR to clipboard, in Go (scrot + tesseract)"
arch=('x86_64')
url="https://github.com/Szmelc-INC/OCR"
license=('MIT')
depends=('tesseract' 'scrot' 'xclip')
makedepends=('git' 'go')
provides=('ocr')
conflicts=('ocr')
options=('!debug')
source=("git+$url.git")
sha256sums=('SKIP')

pkgver() {
  cd "$_repo"
  # ignore descriptive tags like "Linux-Hypr" — plain rev-count reads better
  printf "r%s.%s" "$(git rev-list --count HEAD)" "$(git rev-parse --short=7 HEAD)"
}

prepare() {
  cd "$_repo"
  export GOPATH="$srcdir/gopath" GOFLAGS="-modcacherw"
  # repo ships no go.mod yet — bootstrap a minimal module for the build
  [ -f go.mod ] || go mod init github.com/szmelc-inc/ocr
  go mod tidy
  go mod download
}

build() {
  cd "$_repo"
  export GOPATH="$srcdir/gopath" CGO_ENABLED=1 GOFLAGS="-trimpath -mod=readonly -modcacherw"
  go build -ldflags "-s -w" -o ocr .
}

check() {
  cd "$_repo"
  export GOPATH="$srcdir/gopath" CGO_ENABLED=1 GOFLAGS="-mod=readonly -modcacherw"
  HOME="$srcdir" go vet .
}

package() {
  cd "$_repo"
  install -Dm755 ocr "$pkgdir/usr/bin/ocr"
  install -Dm644 README.md -t "$pkgdir/usr/share/doc/$pkgname/"
  install -Dm644 LICENSE "$pkgdir/usr/share/licenses/$pkgname/LICENSE"
}
