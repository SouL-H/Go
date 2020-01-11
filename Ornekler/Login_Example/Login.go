// example
package main

import (
	"fmt"
)

func main() {
	isim := isim_al()
	yas := yas_al()
	fmt.Printf("\nHoşgeldiniz %s ", isim)
	kontrol_et(yas)

}
func isim_al() (isim string) {
	fmt.Print("İsminiz:")
	fmt.Scanln(&isim)
	return
}
func yas_al() (yas int) {
	fmt.Print("Yaşınız:")
	fmt.Scanln(&yas)
	return
}

func kontrol_et(yas int) {
	switch {
	case yas < 18:
		fmt.Println("Sisteme giriş için yaşınız müsait değil.")
	case yas < 100:
		fmt.Println("\nSisteme başarılı olarak giriş yaptınız.")
		fmt.Printf("Yaşınız= %d\n ", yas)
		oturum_ac()
	default:
		fmt.Println("Yaşınızı kontrol ediniz.")
	}
}
func oturum_ac() {
	fmt.Println("\nKeyifli vakitler.")
}
