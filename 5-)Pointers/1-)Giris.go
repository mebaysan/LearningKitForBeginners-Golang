package main

import (
	"fmt"
	"reflect"
)

/*
Pointer yani işaretçi; bir değişkenin bellek adresini tutan bir başka değişkendir.
Fonksiyonlarla beraber kullanılırken, 'veriler' yerine 'verilerin adresini' gönderebiliriz.

İşaretçilerle çalışırken '&' ve '*' simgelerini kullanırız.
-) '&' simgesi değişkenin adresini temsil eder
-) '*' simgesi işaretçinin gösterdiği adresteki veriyi temsil eder
*/
func main() {
	sayi := 10
	adres := &sayi // sayi değişkeninin adresini adres adlı değişkene atadık
	fmt.Println("Sayı -> ", sayi)
	fmt.Println("Adres -> ", &sayi)
	fmt.Println("Adresteki Veri -> ", *adres)
	fmt.Println("adres değişkeninin tipi -> ", reflect.TypeOf(adres))
	fmt.Println("---------------")
	// istersek 'new' fonksiyonu ile de bir işaretçi oluşturabiliriz. parametre olarak istenilen işaretçi tipi verilir
	adres = new(int) // adres adında bir işaretçi oluşturduk
	*adres = 2020    // adres işaretçisinin temsil ettiği veriyi set ettik
	fmt.Println("adres ->", adres)
	fmt.Println("adresteki veri ->", *adres)
	// Go'da garbage-collector vardır. Bu sayede işi biten işaretçi(pointer)leri bellekten temizleyecektir.
	fmt.Println("---------------")

	x := 0
	fmt.Println(fonk(&x))
}

func fonk(x *int) int {
	return *x + 1
}
