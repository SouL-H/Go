package main

import (
	"bytes"
	"fmt"
)

func main() {
	var x bytes.Buffer
	for i := 0; i < 10; i++ {
		x.WriteString(rastgeleString()) //buradan gelen veriyi yukarıdaki buffere string ekle.
	}
	fmt.Println(x.String())
}

func rastgeleString() string {
	return "xys21314-"
}
