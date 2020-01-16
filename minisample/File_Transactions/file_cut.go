package main

import (
	"log"
	"os"
)

func main() {

	err := os.Truncate("demo.txt", 100) //dosya önceden oluşturulmuş düşünülsün.100 kelime kalana kadar kesti dosyayı.
	if err != nil {
		log.Fatal(err)
	}
}
