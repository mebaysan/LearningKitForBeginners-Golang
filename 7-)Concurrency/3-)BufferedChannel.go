package main

import (
	"fmt"
	"time"
)

func main() {

	kanal := make(chan int , 1) // tamponlu kanal tanımladık -> 1 adet veri/mesaj tutabileceğini tanımladık. FIFO (first in first out)  mantığı ile çalışmaktadır (ilk giren ilk çıkar)
	go kanala_yaz(kanal)
	time.Sleep(time.Millisecond * 100) // 5 saniye bekliyoruz

}

func kanala_yaz(kanal chan int) {
	for i := 0; i < 20; i++ { 
		kanal <- i
		kanaldan_al(kanal)
		i++
	}
	close(kanal) // kanalı kapatabiliriz
}

func kanaldan_al(kanal chan int) {
	toplam := 0
	toplam += <-kanal
	fmt.Println("Toplam -> ", toplam)
}
