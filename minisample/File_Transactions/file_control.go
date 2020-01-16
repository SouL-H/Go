package main

import (
	"log"
	"os"
)

//Dosya var mı yok mu kontrol.

var (
	err      error
	fileInfo *os.FileInfo
)

func main() {
	fileInfo, err := os.Stat("demo.txt")
	if err != nil {
		if os.IsNotExist(err) { //dosya yoksa hata varsa aşağıdaki uyarı yazacak.
			log.Fatal("File does not exist.")
		}
	}
	log.Println("File does exist.File information:")
	log.Println(fileInfo)

}
