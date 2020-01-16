package main

import (
	"log"
	"os"
)

func main() {

	/*

		//V1
		//Salt okunur açılsın.

		file , err := os.Open("demo.txt")
		if err !=nil{
			log.Fatal(err)
		}
		file.Close()
	*/

	//V2
	//OpenFile çok seçenekli dosya açma sistemi

	file, err := os.OpenFile("demo.txt", os.O_CREATE, 0666) //yapılacak işlem ve erişim kodu veriliyor.
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	//Bilgi

	/*
		os.O_RDONLY : Sadece okuma.
		os.O_WRONLY : Sadece yazma.
		os.O_RDWR : Okuma ve yazma yapabilir.
		os.O_APPEND : Dosyanın sonuna ekle.
		os.O_CREATE : Dosya yoksa oluştur.
		os.O_TRUNK : Açılırken dosyayı kes.
		Birden fazla kullanmak için
		os.O_RDONLY|os.O_CREATE|os.O_APPEND bu tarz kullanım yapılır.
		//Cihan Özcana teşekkürler :)
	*/
}
