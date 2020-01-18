package main

import (
	"io"
	"log"
	"os"
)

// Vereceğimiz kadar boyutta byte okumak.
func main() {
	file, err := os.Open("demo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := make([]byte, 15)                 //okunacak veri uzunluğunu belirttik 20 olarak.
	byteRead, err := io.ReadFull(file, byteSlice) //V1den farkı budur paket değiştirdik.
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d bayt okundu.", byteRead)
	log.Printf("Okunan veri: %s", byteSlice)
}
