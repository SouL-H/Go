package main

import (
	"fmt"
	"log"
)

func main() {
	/*
		var s = make([]int, 3)
		n := copy(s, []int{1, 2, 3, 4})

		fmt.Println(n) //üç kayıt diğer tarafa aktarıldı.

		fmt.Printf("Sonuc v1: %v", s) // bu şekilde kopyalandı kapasitesinden dolayı 3 dedik 4 olsa tamamı gider.
	*/
	/*
		s := []int{0, 1, 2, 3}
		_ = copy(s, s[1:])         //0 hariç n e aktarır.
		fmt.Printf("Sonuc: %v", s) //çıktı 1-2-3-3 olmasının sebebi
		// aynı s değerine 0-1-2-3 1.hariç 2-3-4 kopyala dedik 1-2-3 yazdı sonra dördüncü değer zaten 3 oda basıldı 1-2-3-3 çıktısı elde edildi.
	*/
	/*
		var x = make([]byte, 5)
		str := "Hello Mars!"
		copy(x, str)
		fmt.Printf("Sonuc : %v\n", x)
		fmt.Printf("Sonuc : %s ", string(x))
	*/

	ints1 := []int{1, 2, 3, 4, 5}
	ints2 := []int{10, 20, 30, 40, 50}

	log.Printf("ints1: %v", ints1)
	log.Printf("ints2: %v", ints2)

	copied := copy(ints1[1:3], ints2)
	fmt.Println(copied) //parçalı güncelleme işlemi yaptık

	log.Printf("ints1: %v", ints1)
	log.Printf("ints2: %v", ints2)

}
