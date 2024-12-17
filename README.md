# OCR
> ### Optical Character Recognition in Go 
<img src="https://i.imgur.com/lHbYss3.png" width="900px">

# Building
> Install required dependencies
```golang scrot tesseract-ocr tesseract-ocr-eng``` \
> And make sure you'r go installation is on $PATH, eg do `export PATH=$PATH:/usr/local/go/bin`

> Run this script to make it quick and easy!
```bash
git clone https://github.com/Szmelc-INC/OCR/ && cd OCR

# Compiling binary
go mod init ocr
go mod tidy
go build -o ocr main.go

# Adding binary to /bin/
sudo chmod +x ocr && sudo mv ocr /bin/ocr

# Cleaning
cd ../ && rm -fr OCR
echo "Complete!" && exit
```
