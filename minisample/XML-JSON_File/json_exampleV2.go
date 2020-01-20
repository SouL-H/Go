package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//Json işlemleri

//Type Nesneleri

// Name structları

type Name struct {
	Family   string
	Personel string
}

type Email struct {
	ID      int
	Kind    string
	Address string
}

type Interest struct {
	ID   int
	Name string
}

type Person struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	Gender    string
	Name      Name
	Email     []Email
	Interest  []Interest
}

func GetPerson(p *Person) string {

	return p.FirstName + " " + p.LastName
}

func GetPersonEmailAdress(p *Person, i int) string {
	return p.Email[i].Address //Hangi personelin nesne örneğini çağırıyorsa ona ait emailler içerisinde o indexe sahip email dönecek.
}

func GetPersonEmail(p *Person, i int) Email { //geriye email nesnesi dönecek.
	return p.Email[i]
}

func WriteMessage(msg string) {

	fmt.Println(msg)
}

func WriteStarLine() {
	fmt.Println("***************")
}
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error : ", err.Error())
		os.Exit(1)
	}
}

func SaveJSON(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)
	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}

func main() {
	person := Person{
		ID:        9,
		FirstName: "Ahmet",
		LastName:  "Yekta",
		UserName:  "AhmetYkt",
		Gender:    "true",
		Name:      Name{Family: "test", Personel: "Ahmet"},
		Email: []Email{
			Email{ID: 1, Kind: "Work", Address: "ahmet@text.com"},
			Email{ID: 2, Kind: "Hobby", Address: "ahmethoby@gmail.com"},
		},
		Interest: []Interest{
			Interest{ID: 1, Name: "Golang <3"},
			Interest{ID: 2, Name: "Python"},
			Interest{ID: 3, Name: "Kafka"},
		},
	}
	WriteMessage("Reading Operation Started")

	WriteMessage("Personel FullName")
	WriteStarLine()
	res := GetPerson(&person) //personelin hafızadaki adresini yolladık.
	WriteMessage(res)
	WriteStarLine()
	WriteMessage("\n")

	WriteMessage("Personel Email With Index")
	WriteStarLine()
	resEmail := GetPersonEmailAdress(&person, 1)
	WriteMessage(resEmail) //Dışardan string alabiliyor.

	WriteStarLine()
	WriteMessage("\n")

	WriteMessage("Personel Email Object With Index")
	WriteStarLine()

	resEmail2 := GetPersonEmailAdress(&person, 0)
	fmt.Println(resEmail2) //Dışarıdan interface tipinde (farklı türler) alabiliyor.
	WriteStarLine()

	WriteMessage("Reading Operation Ended")
	WriteMessage("\n")

	WriteMessage("Writing Operation Started")
	SaveJSON("person.json", person)
	WriteMessage("Writing Operation Ended")

}
