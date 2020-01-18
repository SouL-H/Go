package main

import (
	"log"
	"os"
)

// Bir dosyaya bayt veri yazma

func main() {

	rawData := "Bu bir test datasıdır."
	newLine := "\n"
	file, err := os.OpenFile(
		"yeni.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	data := rawData + newLine
	byteSlice := []byte(data)
	byteWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d bayt yazıldı \n", byteWritten)
}
