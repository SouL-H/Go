package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := `{

			"data": {
				"object":"card",
				"id":"card_745484564654",
				"first_name":"Murat",
				"last_name":"Duyar",
				"balance":"4.345"
			} 
		}`

	var m map[string]map[string]interface{}                     //birbiri içerisinde iç içe map oluşturuyoruz.
	if err := json.Unmarshal([]byte(jsonStr), &m); err != nil { //m nesnesine aktaracağız.str verisini byte dizisine dönüştürüyoruz.
		//unmarshal jsondan go objelerine veri atkarmaya dönüştürmeye sağlar.
		panic(err)
	}

	fmt.Println(m)

	b, err := json.Marshal(m)
	//Unmarshal tam tersi elimizde m olan go nesnesini jsona dönüştürülmesini sağlayan fonksiyondur.
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b)) //tanımlanan b yi stringe dönüştürüp ekrana yazıyoruz.

}
