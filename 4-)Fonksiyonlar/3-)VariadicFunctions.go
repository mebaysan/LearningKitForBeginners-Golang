package main

import "fmt"

// n sayıda parametre alma, python -> *args gibi
// parametrelerin alındığı yerde 'parametre_adi ...tip' gibi bir kullanımı vardır -> 'sayilar ...int' gibi
func main() {

	toplam := topla(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Toplam -> ", toplam)
}

func topla(sayilar ...int) int {
	var toplam int

	for _, sayi := range sayilar {
		toplam += sayi

	}

	return toplam
}
