package main

import "fmt"

func main() {
	sayi1 := 3
	sayi2 := 5
	toplam := topla(sayi1, sayi2)
	fmt.Printf("sayi1 + sayi2 -> %d\n\n", toplam)

	notlar := []int{100, 34, 56, 78, 94, 23}
	ortalama := ortalama_al(notlar)
	fmt.Printf("notların ortalaması -> %d\n", ortalama)
}

/*
Fonksiyonlar tanımlanırken 'func fonksiyon_adi(){}' şeklinde tanımlanır
*/

func topla(a, b int) int { // fonksiyonun hangi veri tipini return edeceği tanımlanırken belirtilmiştir
	toplam := a + b
	return toplam
}

func ortalama_al(elemanlar []int) int { // elemanlar adında int tipinde bir slice alacak
	var toplam int
	for _, puan := range elemanlar { // range elemanlar -> elemanlar slice'ı dön her döndüğün eleman puan olsun
		toplam += puan
	}
	return toplam / len(elemanlar)
}
