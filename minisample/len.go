package main

import "log"

func main() {


	slice := make([]byte, 10)
	log.Printf("slice: %d",len(slice))//içerisindeki elamanların sayısını belirtiyor.

	m := make(map[string]int)

	m["hello"]=1
	log.Printf("map : %d",len(m))

	channel :=make(chan int , 5)
	log.Printf("kanal : %d", len(channel))
	channel <- 1 //deger atıyoruz
	log.Printf("kanal: %d", len(channel))

	var pointer *[5]byte
	log.Printf("pointer : %d", len(pointer))
}