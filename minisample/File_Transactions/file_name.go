package main

import (
	"log"
	"os"
)

func main() {
	originalPath := "text.txt"
	newPath := "newname.txt" //taşıma işlemi kopyalamada yapılabilir.Şuan isim değiştireceğiz.
	err := os.Rename(originalPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}
