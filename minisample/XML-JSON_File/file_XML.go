package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("sites.xml")
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Error:  %v", err)
		return
	}

	v := ObjectSites{}            //Komple xml kullanacağız.
	err = xml.Unmarshal(data, &v) //Datayı veriyoruz ve nesneyi al unmarshal işlemi yap.
	if err != nil {
		fmt.Printf("Error:  %v", err)
		return
	}

	fmt.Println(v)
}

type ObjectSites struct {
	XMLName     xml.Name `xml:"sites"`
	Version     string   `xml:"version,attr"`
	sites       []site   `xml:"site"`
	Description string   `xml:",innerxml"`
}

type site struct {
	XMLName     xml.Name `xml:"site"`
	Name        string   `xml:"Name"`
	Description string   `xml:"Description"` //bunları gördüğünde eşleştir anlamı katmak için bu yapıyı oluşturduk.
	Category    string   `xml:"Category"`
}
