# ***OCR*** [v2]
## [Wayland / Hyprland]
> Optical Character Recognition in Golang, made for Linux (X11 + Wayland) 
<img src="https://i.imgur.com/5Dgktbd.png" width="900px">

# Setup
> Run `setup.sh` (NOT TESTED FULLY YET!) \
> And make sure you'r go installation is on $PATH, eg do `export PATH=$PATH:/usr/local/go/bin`

## > Get started with:!
```bash 
git clone https://github.com/Szmelc-INC/OCR.git && cd OCR

# Setup script
chmod +x setup.śh && bash setup.sh

# Adding binary to $PATH (/bin/)
sudo chmod +x ocr && sudo mv ocr /bin/ocr

# cleanup
cd ../ && rm -fr OCR
echo "Complete!" && exit
```

> Manual
```bash
# Compiling binary
go mod init ocr
go mod tidy
# optional step to get deps from repos (see setup.sh)
go build -o ocr main.go
```
