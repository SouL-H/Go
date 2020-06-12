package main


//Mevcut kütüphaneler
import (
	"fmt"
	"math"
)

func main() {
	//Adım Sayısı kullanıcıdan input alınır.
	var adimsayisi int
	fmt.Printf("Adim sayisi giriniz:")
	fmt.Scanf("%d", &adimsayisi)
	fmt.Print(adimsayisi)
	//Çarpımsal Yöntem ile oluşturulan random sayı dizisi
	x := 53
	m := 100
	a := 37
	dizi := [1000]int{}
	floatdizi := [999]float64{}
	dizi[0] = (a * x) % m
	for i := 0; i < 999; i++ {
		dizi[i+1] = (dizi[i] * a) % m
		floatdizi[i] = float64(dizi[i]) / 100
	}

	//Monte Carlo Yöntemi ile Pi (π) Sayısının Bulma
	fmt.Print("\n")
	sayac := 0
	randomsayaci := 0
	for i := 0; i < adimsayisi; i++ {
		randomsayaci++

		//Diziden unique sayı bulmak için ürettiğim denklem
		p := floatdizi[(randomsayaci%10)+((randomsayaci*adimsayisi)%100/100)]
		q := floatdizi[(randomsayaci*adimsayisi%100/10)+((randomsayaci+8*adimsayisi)%100/100)]

		//fmt.Print(q, p, "\n") /* Mevcut üretilen random sayı kontrolü */

		if p*p+q*q < 1.0 {
			sayac++
		}
	}
	benimPI := float32(4*sayac) / float32(adimsayisi)
	hata := float64(100.0*math.Abs(float64(math.Pi-benimPI))) / float64(benimPI)
	fmt.Print("Mevcut PI degeri:", math.Pi)
	fmt.Print("\nMonte Carlo Yöntemi ile bulunan PI degeri:", benimPI, "\n")
	fmt.Print("Hata Payi:", hata, "\n")
}
