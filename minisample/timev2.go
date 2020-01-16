package main

import (
	"fmt"
	"time"
)

func main() {
	x := fmt.Println // ekrana çıktı almaada kısa yol için bunu yaptık.
	xTime := time.Date(1071, 10, 11, 20, 23, 0, 0, time.UTC)
	now := time.Now()
	x(now)
	/*
		//Mevcut değerleri alma
			x(now.Year())
			x(now.Month())
			x(now.Day())
			x(now.Hour())
			x(now.Minute())
			x(now.Second())
			x(now.Nanosecond())
			x(now.Location())
			x(now.Weekday())

	*/
	/*
		//Tarih kontrolü
			x(xTime.Before(now)) //x tarihinden önce mı sorusu sorduk
			x(xTime.After(now))  // x tarihinden sonra mi sorusunu sorduk
			x(xTime.Equal(now))  // x tarihine eşit mi
								//sonuçlar bool değer dönecektir.
	*/
	diff := now.Sub(xTime) //Şimdiki zamanla xTime arasındaki farkı yazdırır.
	x(diff)                //SANİYE
	x(diff.Hours())        //SAAT
}
