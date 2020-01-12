package main

import (
	"fmt"
)

//stringi byte dönüştürme dosya network programlamada
//int valueyi bir stringe dönüştürmede
func main() {

	/*
		data := []byte("Bu bir veri")
		fmt.Println(data)

		var myNo uint64

		myNo = 1132165165121321321

		str := strconv.FormatUint(myNo, 10)
		// fmt.Println("No: " + myNo) dönüşümsüz hatalı olacaktır.
		fmt.Println("No: " + str)

		x := 69
		s := string(x)
		fmt.Println(s) //char gibi algılayıp string değer bastı.
	*/
	var str string = "Bu bir yazi"
	fmt.Println(str)

	var b []byte
	b = []byte(str)
	fmt.Println(b) //stringi byte dönüştürüp ve byte olarak ekrana yazdık

	for x := range b {
		fmt.Println(string(b[x])) //dizideki her byteyi str dönüştürüp ekrana yazıyoruz burada
	}

	str = string(b)
	fmt.Println(str)

}
