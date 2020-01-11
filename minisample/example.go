package main


import "log"

func main(){

	//v1

	/*
		states := make(map[string]string)//şekil bilgileri tutulacak.nesnenin genel yapısı
	fmt.Println("Boş: ", states)

	states["IST"]="Istanbul"
	states["NYC"]="New York"
	states["LND"]="London"
	fmt.Println("Dolu : ", states)

	lnd :=states["LND"]
	fmt.Println("London: " , lnd)

	delete(states , "NYC")
	fmt.Println("Dolu:", states)
	*/
/*
	//v2

	unbuffered := make(chan int)
	log.Printf("unbuffered: %v,type: %T , len: %d , cap: %d", unbuffered,unbuffered,len(unbuffered),cap(unbuffered)) //kullanılan hafıza bilgilerini yazma
	//v3
	buffered := make(chan int,10)
	log.Printf("buffered: %v,type: %T , len: %d , cap: %d", buffered,buffered,len(buffered),cap(buffered))
*/

//Şimdi byte dizisi üzerinde yapılsın.

	slice := make([]byte,7)
	log.Printf("slice: %v,type: %T , len: %d , cap: %d", slice,slice,len(slice),cap(slice))
	slice1 := make([]byte,4,9)//kapasitesi belli sadece 4 tanesini kullanabiliyoruz.Kapasite aktif ama :)
	log.Printf("slice1: %v,type: %T , len: %d , cap: %d", slice1,slice1,len(slice1),cap(slice1))
}