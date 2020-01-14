package main

import (
	"fmt"
	"strings"
)

func main() {

	//v1

	/*
		value1 := "Biz çok zenginiz"

		fmt.Println(value1)

		result1 := strings.Replace(value1, "zenginiz", "zengin miyiz", -1) //stringde kelime değiştirmeyi sağlar. -1 sonuna kadar git anlamına gelmektedir.
		fmt.Println(result1)

		value2 := "insan insan insan mıdır"

		result2 := strings.Replace(value2, "insan", "hayvan", 2) //ilk iki tane insanı hayvan olarak çevir diyoruz.
		fmt.Println(result2)
	*/

	//v2
	/*
		value := "kuş ,kedi ve balık"

		r := strings.NewReplacer("kuş", "köpek", "kedi", "aslan", "balık", "maymun") //buradaki parametre kuş yerine köpek kedi yerine aslan yaz demek anlamına geliyor.
		//bu şekilde verdiğimiz değerdeki veriler karşılık gelen verilerle değiştirildi.
		result := r.Replace(value)
		fmt.Println("Sonuç :" + result)
	*/

	//v3

	x := "Bu 'bir' açıklama satırı"

	x = strings.Replace(x, "'", "\\", -1)

	fmt.Println(x)
}
