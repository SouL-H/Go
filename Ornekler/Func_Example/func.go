// func
package main

import (
	"fmt"
)

func main() {
	aday := []int{40, 100, 15, 70, 45, 100, 54, 78}
	fmt.Println("Ortalama=", ort(aday))
}
func ort(elemanlar []int) int {
	var toplam int
	for _, puan := range elemanlar {
		toplam += puan
	}

	return toplam / len(elemanlar)

}
