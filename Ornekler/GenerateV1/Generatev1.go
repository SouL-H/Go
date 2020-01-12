package main

import (
	"fmt"
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().UnixNano()) // rastgele veri üretmemizi sağlayan nesne

const _charset = "abcdeghijklmnopqrstuvwxyzABCDEGHIJKLMNOPQRSTUVWXYZ123456789" //bunları kullanarak random değer üretilicek.

func GeneratePassword(length int) string {
	x := make([]byte, length)
	for i := range x {
		x[i] = _charset[source.Int63()%int64(len(_charset))] //genellikle kullanılan algoritma charset içindeki.
	}
	return string(x)
}
func main() {

	password := GeneratePassword(15)
	fmt.Println(password)
}
