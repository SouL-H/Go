package main

import "fmt"

func main() {

	/*
		str := "sad sad
		sadasd
		asdasd"
	*/
	message := `sad sad
	sadasd
	asdasd` // doğru kullanım alt+9+6 klavye kombinasyonu ile tek tırnak oluşturulur.

	fmt.Println(message)
	fmt.Printf("%s", message) // bu şekilde de çıktı alınabilir.
}
