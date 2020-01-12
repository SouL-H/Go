package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	t0 := time.Now()
	// V1 builder ile birleştirme.
	for i := 0; i < 100000; i++ {
		builder := strings.Builder{}
		for v := 0; v < 100; v++ {
			builder.WriteString("Plüton")
		}
		if builder.Len() == 0 {
			break
		}
	}

	t1 := time.Now()
	for i := 0; i < 100000; i++ {
		result := ""
		for v := 0; v < 100; v++ {
			result += "Plü"
		}
		if len(result) == 0 {
			break
		}
	}
	t2 := time.Now()
	fmt.Println("V1 : ", t1.Sub(t0)) //sub komutu t1-t0 arasındaki farkı al anlamına geliyor.
	fmt.Println("V2 : ", t2.Sub(t1))
}
