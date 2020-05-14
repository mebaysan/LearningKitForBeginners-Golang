package main

import (
	"fmt"
)

// Go'da tek döngü vardır. FOR

// FOR döngüsünde başlangıç kuralı (i:=0) ve ilerleme kuralı  (i++) pas geçilebilir

/*
for ifade1; ifade2; ifade3 {


}

*/

func main() {
	var sayac = 1
	for i := 0; i <= 4; i++ {
		fmt.Println(sayac, ". sayı = ", i)
		sayac++
	}

	// FOR'un while gibi kullanılması

	var sayi int = 1
	for sayi <= 5 { // başlangıç değeri (i:=1) ve ilerleme kuralı (i++) pas geçilebilir
		fmt.Println("While döngüsü gibi düşün, sayı = ", sayi)
		sayi++
	}

 	// go da while true
	// var deneme bool = true

	// for deneme{
	// 	fmt.Println("Naber")
	// }

	

}
