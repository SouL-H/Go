package main

import (
	"log"
)

func main() {

	m := make(map[string]int)
	log.Println(m)

	m["01"] = 1
	log.Println(m)

	m["02"] = 2
	log.Println(m)

	m["03"] = 3
	log.Println(m)

	delete(m, "02") //02 keyini siliyoruz.
	log.Println(m)

	m = nil         // m sıfırlamış oluyoruz null eşitlemiş oluyoruz.
	delete(m, "01") // silsekte değişen bir  şey olmayacak boş map yani.

}
