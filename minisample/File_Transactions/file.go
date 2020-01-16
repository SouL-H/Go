package main

import (
	"log"
	"os"
)

var ( //global tanımlıyoruz.Farklı yerlerde kullanmak için.
	newFile *os.File
	err     error
)

func main() {
	newFile, err = os.Create("Demo.txt") //dosya oluşturur.
	if err != nil {
		log.Fatal(err)
	}
}
