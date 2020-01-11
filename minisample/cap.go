package main

import "log"

func main() {
	slice := make([]byte, 2, 5)
	log.Printf("slice :%d", cap(slice)) //kapasite
	log.Printf("slice :%d", len(slice)) //uzunluğu
	log.Printf("slice :%v", slice)      //içindekiler %v varsayılan değeri gösterirmiş.

	channel := make(chan int, 10)
	log.Printf("channel :%d", cap(channel))
	log.Printf("channel :%d", len(channel))

	var pointer *[15]byte
	log.Printf("pointer: %d == %d", cap(pointer), len(pointer))

}
