package main

import (
	"fmt"
)

// Eş zamanlı olarak çalışan işlemler arasında birbirleriyle iletişim kurabilmeleri için 'kanal' adı verilen (channels) yöntem kullanılır

func main() {
	// var kanal1 chan string = make(chan string) // string türünde verileri taşıyacak bir channel oluşturulmuştur
	kanal1 := make(chan string) // string türünde verileri taşıyacak bir channel oluşturulmuştur

	go func() {
		kanal1 <- "Merhaba" // kanal1'e    "Merhaba" stringi koyuluyor
	}() // anonim bir fonksiyon oluşturduk. channels'lar ile haberleşme yapabilme kabiliyeti sadece goroutine'ler ile beraber yapılabilir. Anonim fonksiyonlar tanımlandığı yerde çalışırlar.

	mesaj := <-kanal1 // kanal1'deki veriler mesaj değişkenine aktarılıyor

	fmt.Println(mesaj)
	kanal2 := make(chan string)
	go aktarim(kanal2) // parametre olarak aldığımız kanalı verdik ve func çalıştırdık
	yeni_gelen := <-kanal2 // aşağıdaki func kanalımıza bir mesaj koymuştu, bu mesajı bir değişkene atadık
	fmt.Println(yeni_gelen)
}

func aktarim(kanal chan string) { // parametre olarak bir channel alacak dedik
	kanal <- "Bu bir mesajdır..." // aldığı channel'a bir string yükleyecek
}
