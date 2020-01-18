package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os" //operation system
)

//Geçiçi dosya ve klasör üstünde çalışacağız.Sonra sileceğiz.

func main() {

	tempDirPath, err := ioutil.TempDir("", "geciciklasör")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Geçiçi klasör oluşturuldu: ", tempDirPath)
	//Geçiçi dosya

	tempFile, err := ioutil.TempFile(tempDirPath, "gecicidosya.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Geçiçi dosya oluşturuldu: ", tempFile.Name())

	// Kapatma işlemleri

	err = tempFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	//Silme

	err = os.Remove(tempFile.Name())
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s dosyası silindi.", tempFile.Name())
	}

	err = os.Remove(tempDirPath)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("\n%s klasörü silindi.", tempDirPath)
	}

}
