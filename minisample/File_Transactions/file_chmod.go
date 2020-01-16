package main

import (
	"log"
	"os"
)

func main(){

	file ,err := os.OpenFile("demo.txt",os.O_WRONLY,0666)//dosya izinlerine bağlı eğer salt okunur olursa alttaki hata yazar ekrana.O_RDONLY olursa okuma izni kontrol edilir.
	iff err != nil{
		if os.IsPermission(err){
			log.Println("Hata:Yazma izni reddedildi.")
		}
	}
	file.Close()
}