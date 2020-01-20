package main

//CSV : Comma-Seperated Values
import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	f, err := os.Open("data.csv")
	checkError(err)

	r := csv.NewReader(f)
	r.Read() //Kayıtları oku

	records, err := r.ReadAll() //Kayıtları elde ediyoruz.
	checkError(err)

	for _, row := range records {
		//Çok boyutlu dizileri okuyoruz
		printRow(row)
	}
}

func printRow(row []string) {
	log.Printf("len(row) %d\n", len(row)) //Satırları toplama işlemi
	for i, col := range row {
		log.Printf("[%d]: %s\n", i, col) //String ve nümerik değer gelecek.
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
