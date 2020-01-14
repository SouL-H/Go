package main

import (
	"fmt"
	"strings"
)

func main() {

	//concains1()
	//contains2()
	//containsany1()
	soon()
}

func concains1() {

	value1 := "test"
	value2 := "golang"

	source := "Bu bir test"

	if strings.Contains(source, value1) { //kaynak içinde veri var mı.
		fmt.Println(1)
	}
	if !strings.Contains(source, value2) { //kaynak içinde golang yok false dönüyor ama ! uyguladık test için yine true yaptık.
		fmt.Println(2)
	}
}

func contains2() {

	fmt.Println(strings.Contains("Turkey", "Tur"))
	fmt.Println(strings.Contains("Human", "HUMAN"))
}

func containsany1() {
	source := "12345"

	if strings.ContainsAny(source, "135") { //135 her hangi biri var mı diye arama yapar
		fmt.Println("X")
	}
	if strings.ContainsAny(source, "543") { //yine her hangi biri varsa true döner
		fmt.Println("Y")
	}
	if strings.ContainsAny(source, "?") { //? olmadığı için ekrana yazmaz
		fmt.Println("Z")
	}
	fmt.Println(strings.ContainsAny("Merhaba", "aio%f")) // herhangi birisi olduğu için true döner.
	fmt.Println(strings.ContainsAny("Mars", "i"))
	fmt.Println(strings.ContainsAny("Turkey", "r & z"))
}

func soon() {

	fmt.Println(strings.ContainsAny("Shell-12541", "1-2")) //tek tek taradığı için true değer döner
	fmt.Println(strings.Contains("Shell-12541", "1-2"))    //komple kalıp olarak taradığı için false
}
