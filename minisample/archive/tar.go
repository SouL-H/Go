package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
)

//Dosyaları TAR dosyası olarak sıkıştırma.
//Tar dosyası üretme
//Go'da fonksiyon isimleri büyük harfle başlarsa public
//Küçük harfle başlarsa private anlamı vardır.

var fileFolderPath = "files\\"
var files = []string{fileFolderPath + "example.go"}

func addFile(fileName string, tw *tar.Writer) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("Dosya açılırken hata meydana geldi. (%s): %s", fileName, err)
	}
	defer file.Close()

	stat, err := file.Stat() //açılan  dosya bilgilerini alacağız.
	if err != nil {
		return fmt.Errorf("Dosya bilgiler alınırken bir hata meydana geldi (%s): %s", fileName, err)
	}

	hdr := &tar.Header{
		ModTime: stat.ModTime(),
		Name:    fileName,
		Size:    stat.Size(),
		Mode:    int64(stat.Mode().Perm()),
	}

	if err := tw.WriteHeader(hdr); err != nil {
		msg := "TAR header yazılırken bir hata meydana geldi. %s: %s"
		return fmt.Errorf(msg, fileName, err)
	}

	copied, err := io.Copy(tw, file) //Headerler ile çalışmak lazım , temel bilgilere erişilebilen işlem yapılabilen alanlardırWriter ile file nesnesi arasındaki değeri kopyalayacağız.
	if err != nil {
		return fmt.Errorf("%s dosyası TAR'a yazılırken hata meydana geldi: %s", fileName, err)
	}

	//arada veri kaybı var mı kontrolü

	if copied < stat.Size() {
		msg := "%s dosyasına %d kadar veri yazıldı , ama beklenen veri %d kadardı."
		return fmt.Errorf(msg, fileName, copied, stat.Size())
	}
	return nil
}
func createArchiveTARFile(archiveFileName string) int {
	if len(archiveFileName) == 0 {
		return -1
	}
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC //Dosyayı yazarken ki izinler
	file, err := os.OpenFile(archiveFileName+".tar", flags, 0644)
	if err != nil {
		log.Fatalf("TAR dosyasına veri yazmak için açılırken hata meydana geldi. : %s", err)
		return -1
	}
	defer file.Close()

	tw := tar.NewWriter(file) //Artık elimizde bir dosya var
	defer tw.Close()

	for _, fileName := range files { //her döngüde dosyanın birini al ekle
		if err := addFile(fileName, tw); err != nil {
			log.Fatalf("%s dosyası TAR dosyasına eklenirken hata meydana geldi : %s", fileName, err)
		}
	}
	return 1 // işlem başarılı olduğunu gönderiyoruz.
}

func main() {
	result := createArchiveTARFile("dosya")
	if result > 0 {
		fmt.Println("İşlem başarılı : ", result)
	} else {
		fmt.Println("İşlem başarısız oldu : ", result)
	}
}
