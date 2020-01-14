package main

import (
	"fmt"
	"strings"
)

func main() {
	// Count ile string içindeki aramalar yapıyoruz.Kaç adet olduğunu bulabiliriz.

	fmt.Println(strings.Count("Moon", "o"))
	fmt.Println(strings.Count("test", ""))   // karakter sayar ve +1 fazla sonuç verir.
	fmt.Println(strings.Count("Ahmet", "a")) //Büyük küçük harfe duyarlıdır.
}
