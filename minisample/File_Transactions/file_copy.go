package main

import (
	"io"
	"log"
	"os"
)

//Dosya kopyalama

func main() {
	fileName := "demo.txt"
	originalFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()

	// Yeni dosya oluşturma

	newFile, err := os.Create("demo_copy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	//Kaynaktan hedefe doğru kopyalama.

	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Baytlar kopyalandı : %d", bytesWritten)
	//Dosya işleme ve kaynakları hafıdan boşatma alanı
	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
}
