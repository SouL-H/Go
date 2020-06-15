// test
package main

import (
	"fmt"
)

func obeb(a int, b int) int {
	if b == 0 {
		return a
	} else {
		return obeb(b, a%b)
	}
}
func main() {
	var giris int
	var tutucu [3]int
	for i := 0; i <= 2; i++ {
		fmt.Print(i+1, ".sayiyi giriniz: \n")
		fmt.Scan(&giris)
		tutucu[i] = giris
	}
	fmt.Print(tutucu)
	fmt.Println(obeb(tutucu[0], obeb(tutucu[1], tutucu[2])))
	/*
		a := 50
		b := 25
		c := 20
		sum := 0

		for i := 1; i <= a && i <= b && i <= c; i++ {
			sum += i
			if a%i == 0 && b%i == 0 && c%i == 0 {
				fmt.Println(i)
			}
		}*/
}
