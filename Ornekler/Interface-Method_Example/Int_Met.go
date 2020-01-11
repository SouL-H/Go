// Murat hocanın 7.bölümün sonunda hata bulmayla ilgili uygulamasının çözümü.
package main

import "fmt"

type Köpek struct {
	tür string
	ses string
}

func (k Köpek) SesCikarma() string {
	return k.ses // buradaki println komutunu kaldırıp string tipinde k.ses'i dönmesini sağladık.
}

type Varlik interface {
	SesCikarma() string //interface'de dönüş değerinin türü olması lazım String türünü belirledik.
}

func SesCikar(x Varlik) {
	x.SesCikarma()
}
func main() {
	var Karabaş Köpek
	Karabaş = Köpek{tür: "Kangal", ses: "Hav Hav!"}
	SesCikar(Karabaş)
	Karabaş.SesCikarma()
	fmt.Println(Karabaş.SesCikarma())
}
