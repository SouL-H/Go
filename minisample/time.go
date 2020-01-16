package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Date(2019, time.November, 10, 20, 0, 0, 0, time.UTC) // zaman bilgisi oluşturuldu.
	fmt.Printf("Çalışıyor: %s\n", t)

	now := time.Now()
	fmt.Printf("Mevcut zaman : %s \n", now) //mevcut zaman

	fmt.Println("Ay:", now.Month())
	fmt.Println("Gün:", now.Day())
	fmt.Println("Haftanın Günü:", now.Weekday())

	tomorrow := now.AddDate(0, 0, 1) //1 gün ekle
	fmt.Printf("Tomorrow is  %v , %v, %v , %v\n", tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())

	longFormat := "Monday,Junuary 2, 2018"

	fmt.Println("Tomorrow is", tomorrow.Format(longFormat)) //Burada bir format belirttik.Derleyici yukardaki formatı yaparak yarının bilgisini almamızı sağlıyor.

	shortFormat := "1/2/2020" // bu formatta istediğimizi belirtik her hangi bir tarih değil.
	fmt.Println("Tomorrow is", tomorrow.Format(shortFormat))
}
