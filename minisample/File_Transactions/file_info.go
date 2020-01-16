package main

import (
	"fmt"
	"log"
	"os"
)

//Dosya bilgisi alma

var ( //fileinfo nesne örneği oluşturacağız.

	fileInfo *os.FileInfo //fileinfo pointerini fileinfoya atadık.package File_Transactions
	err      error        //error işlemleri için error nesnesi tanımladık.
)

func main() {

	fileInfo, err := os.Stat("demo.txt") //belirtilen dosya
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes ", fileInfo.Size())
	fmt.Println("Permission: ", fileInfo.Mode())       //dosya dizini
	fmt.Println("Last Modified: ", fileInfo.ModTime()) //son değiştirme tarihi
	fmt.Println("Is Directory: ", fileInfo.IsDir())    // dosyamı dizin mi olduğu
	fmt.Printf("System interface type%T\n: ", fileInfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileInfo.Sys())

}
