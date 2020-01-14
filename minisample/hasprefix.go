package main

import (
	"fmt"
	"strings"
)

func main() {

	data := "Turkey"
	result := strings.HasPrefix(data, "u") //Verilen değerin başlangıç kısında u var mı diye kontrol eder.Küçük ve büyük harfe duyarlı.
	fmt.Println(result)
	result = strings.HasSuffix(data, "key") // Verilen değerin son kısmından arama yapar.
	fmt.Println(result)

}
