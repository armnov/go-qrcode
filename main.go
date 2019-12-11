package main

import (
	qrcode "github.com/skip2/go-qrcode"
	"fmt"
)

func main(){
	err := qrcode.WriteFile("https://example.org", qrcode.Medium, 256, "qr.png")
	if err != nil {
		fmt.Println("Done.")
	} else {
		fmt.Println("Something is wrong.")
	}
	
}
