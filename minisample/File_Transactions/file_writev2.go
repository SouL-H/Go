package main

import (
	"io/ioutil"
	"log"
)

func main() {
	fileName := "yenibir.txt"
	data := "Selam Dünyalı"

	err := ioutil.WriteFile(fileName, []byte(data), 0666) // her hangi bir dosya yoktu oluşturmayı otomatik yapar.
	if err != nil {
		log.Fatal(err)
	}

}
