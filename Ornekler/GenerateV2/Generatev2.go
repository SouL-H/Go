package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Option struct { // seçenekleri tutmamız için nesneye ihtiyacımız var.
	length    int
	upperCase bool
	lowerCase bool
	numbers   bool
	specials  bool
}

var source = rand.NewSource(time.Now().UnixNano())

const _charsetLowercase = "abcdeghijklmnopqrstuvwxyz"
const _charsetUppercase = "ABCDEGHIJKLMNOPQRSTUVWXYZ"
const _charsetNumbers = "0123456789"
const _charsetSpecials = "=-_[.~]!#+&"

func GeneratePassword(opt Option) (string, error) {
	x := make([]byte, opt.length)
	charset := "."

	if opt.upperCase {
		charset += _charsetUppercase
	}
	if opt.lowerCase {
		charset += _charsetLowercase
	}
	if opt.specials {
		charset += _charsetSpecials
	}
	if opt.numbers {
		charset += _charsetNumbers
	}
	if charset == "." {
		return "NON-GENERATED!", errors.New("Her hangi bir karakter seçimi zorunludur.")
	}

	for i := range x {
		x[i] = charset[source.Int63()%int64(len(charset))]

	}

	return string(x), nil

}

func main() {
	password, err := GeneratePassword(Option{length: 20, upperCase: true, lowerCase: true, specials: true, numbers: true})
	fmt.Println(password)
	if err != nil {
		fmt.Println(err.Error())
	}

}
