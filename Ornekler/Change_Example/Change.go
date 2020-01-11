// func
package main

import (
	"fmt"
)

/* func degistir(x, y int) (int, int) { //burada iki tane tam sayı döndüreceğini belirttik
	return y, x
} */

func main() {
	/* a := 1
	b := 2
	fmt.Println(a, b)
	a, b = degistir(a, b)
	fmt.Println(a, b) */
	//şimdi pratik olarak yer değiştirmeyi işleyelim.

	var a, b int = 1, 2
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)
}
