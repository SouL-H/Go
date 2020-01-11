package main

import (
	"fmt"
)

func main() {
	var toplam int
	bilgisayar := map[string]int{
		"Dizüstü":  27,
		"Masaüstü": 123,
		"Tablet":   8,
		"Sunucu":   3,
	}
	for _, deger := range bilgisayar {

		//var toplam int

		/* Kodu aldığım kaynakta hata bulmacasıydı.
		Hatanın sebebi yukardaki koddur.For döngüsü içerisinde tekrardan toplam değişkeni oluşturulduğu için.
		Yukardan ne değer gelirse gelsin int değeri baştan 0 olarak geldiği için sürekli 0 ile toplanmıştır.
		Bu yüzdende sonuç değişkeni hatalı çıkmaktadır.// kaldırarak deneyebilirsiniz.
		*/
		toplam += deger
	}
	fmt.Printf("Toplam %d bilgisayar var.\n", toplam) //Prinf kullanarak aynı C dersindeki gibi %d,%f,%v parametrelerini kullanabiliriz.
}
