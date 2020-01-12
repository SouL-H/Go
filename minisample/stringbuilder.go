package main

import (
	"fmt"
	"strings"
)

//V2 string birleştirme
func main() {

	builder := strings.Builder{} //builder nesne örneği oluşturuyor.

	for i := 0; i < 10; i++ {
		builder.WriteString("Data ") //10 tane data verisi oluyor elimizde.
	}
	result := builder.String() //builder içindeki byteyi stringe dönüştürüyor.
	fmt.Println(result)
	//Diğer yöntemlerden daha performanslıdır bu
}
