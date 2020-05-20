package main

import "fmt"

func main() {
	kanal := make(chan string, 1)
	metin(kanal, "Ses Deneme!!!!...")
	ses_seviyesi(kanal, 4)
}

func ses_seviyesi(kanal <-chan string, ses int) { // ses_seviyesi fonksiyonu yalnızca kanala veri gönderebilir (<-chan)
	fmt.Printf("Anons ses seviyesi: %d\n", ses)
	fmt.Printf("Anons metni: %s\n", <-kanal)
}

func metin(kanal chan<- string, metin string) { // metin fonksiyonu kanaldan yalnızca metin okuyabilecektir (chan<-)
	kanal <- metin
}
