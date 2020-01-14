package main

import (
	"fmt"
	"strings"
)

func main() {

	txtString := "golang"
	repString := strings.Repeat(txtString, 7) // 7 kere üstteki stringi tekrar et anlamı taşır.
	fmt.Println(repString)
	fmt.Println("ba" + strings.Repeat("na", 2)) // farklı kullanım.
}
