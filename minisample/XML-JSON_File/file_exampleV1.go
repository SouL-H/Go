package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//Xml üstünden yönetim
func main() {

	xmlFile, err := os.Open("Employees.xml")
	if err != nil {
		fmt.Println("Opening file error: ", err)
		return
	}

	defer xmlFile.Close()

	xmlData, _ := ioutil.ReadAll(xmlFile)
	var c Company              // Company nesnesi oluşturduk
	xml.Unmarshal(xmlData, &c) //xml verisini tutan bir nesne artık
	//xml data ile c nesnesini veriyoruz.Tek company nesnesi ile tüm çalışan ve şirket bilgisi aktarılır.
	fmt.Println(c.Persons) // Artık o veriyi aldık xml datadan c ye aktardık.

	//Convert to JSON

	var person JsonPerson
	var persons []JsonPerson

	for _, value := range c.Persons {
		person.ID = value.ID
		person.FirstName = value.FirstName
		person.LastName = value.LastName
		person.UserName = value.UserName

		persons = append(persons, person) //her oluşturulan person listesini aktar.
	}
	jsonData, err := json.Marshal(persons) //elimizdeki persons dizisini jsona çevir
	if err != nil {
		fmt.Println(err)
		os.Exit(1) //uygulamayı kapat
	}
	//ekrana yazma

	fmt.Println(string(jsonData))

	//Dosyaya yaz

	jsonFile, err := os.Create("./Employees.json")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer jsonFile.Close() //iş bitince kapat dosyayı
	jsonFile.Write(jsonData)
	jsonFile.Close() //buradaki tüm işlemleri bitir.
}

type JsonPerson struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
}

type Person struct {
	XMLName   xml.Name `xml:"person"`
	ID        int      `xml:"id"`
	FirstName string   `xml:"firstname"`
	LastName  string   `xml:"lastname"`
	UserName  string   `xml:"username"`
}

type Company struct {
	XMLName xml.Name `xml:"company"`
	Persons []Person `xml:"person"`
}

func (p Person) String() string { //person nesnesine bağlı bir string metotu
	return fmt.Sprintf("\t ID: %d - FirstName: %s - LastName: %s - UserName: %s \n", p.ID, p.FirstName, p.LastName, p.UserName)
}
