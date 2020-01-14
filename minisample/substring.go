package main

import (
	"fmt"
)

func main() {

	value := "Merhaba Jupiter"

	runes := []rune(value) //valueyi dizi haline getirdik.

	subStr1 := string(runes[0:4])
	fmt.Println("Substring:", subStr1)

	subStr2 := value[0:6]
	fmt.Println("ASCII Sub:", subStr2)

	sub1 := value[2:len(subStr1)]
	fmt.Println("sub1:", sub1)
}
