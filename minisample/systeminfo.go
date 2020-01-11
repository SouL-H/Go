package main

import (
	"fmt"
	"os"
)

func main(){

	uName := os.Getenv("USERNAME")
/*	uDomain := os.Getenv("USERDOMAIN")
	procossorArchitecture := os.Getenv("PROCESSOR_ARCHITECTURE")
	//processorIdentifies := os.Getenv("PROCESSOR_IDENTIFIER")
	//processorLevel := os.Getenv("PROCESSOR_LEVEL")
	goRoot := os.Getenv("GOROOT")
	goPath := os.Getenv("GOPATH")
	homePath := os.Getenv("HOMEPATH")

	
	fmt.Println("Domain :"+uDomain)
	fmt.Println("İşlemci Mimarisi: " + procossorArchitecture)
	fmt.Println("HOMEPATH :" + homePath)
	fmt.Println("GOPATH :" + goPath)
	fmt.Println("GOROOT : " + goRoot ) */ // Komut bilgileri

	fmt.Println("Username :" + uName)
	/* for _,env := range os.Environ(){
		fmt.Println(env) // ortam değişkenleri
		
	}*/


}