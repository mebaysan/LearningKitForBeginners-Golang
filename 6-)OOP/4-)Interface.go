package main

import (
	"fmt"
)

func main() {

	kus := Kus{}
	fmt.Println(SesCikarma(kus))

}

type Hayvan interface { // bir arayüz oluşturduk
	SesCikar() string // SesCikar adında bir fonksiyonu olacak bu arayüzün
}

func SesCikarma(hayvan Hayvan) string { // arayüzü parametre olarak alacak olan fonksiyon (SesCikarma). Kendisine parametre olarak gelen objenin, yapının 'SesCikar' adındaki fonksiyonunu çalıştıracak. Çünkü Hayvan interface'de 'SesCikar' adında metot var
	return hayvan.SesCikar()
}

type Kus struct {
}

func (self Kus) SesCikar() string {
	return "Cik cik cik...!"
}
