package main

import (
	"fmt"
	"time"
)

// Eş zamanlı olarak çalışmasını istediğimiz fonksiyonun başına 'go' ön eki kullanmamız yeterlidir. Goroutine.

func main() {
	go sayilari_sirala()
	go harfleri_sirala()
	time.Sleep(time.Millisecond * 100) // 1 saniye bekliyoruz
}

func sayilari_sirala() {
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
}

func harfleri_sirala() {
	harfler := "abcdefgğhıijklmnoöprsştuüvyzwx"
	for _, harf := range harfler {
		fmt.Print(string(harf), " ")
	}
}
