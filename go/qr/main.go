package main

import (
	"log"

	"github.com/skip2/go-qrcode"
)

var content = "https://drive.google.com/file/d/1q5H_MgMQVGZ48f1MZKCnvyk6UM7KYElK/view?usp=sharing"

func main() {
	err := qrcode.WriteFile(content, qrcode.Medium, 128, "qr.png")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
}
