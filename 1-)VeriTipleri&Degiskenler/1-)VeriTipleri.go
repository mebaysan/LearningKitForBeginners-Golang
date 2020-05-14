package main

import "fmt"
// GO dilinde değişken tanımlayıp kullanmazsak veya paket çağırıp kullanmazsak hata alırız
func main() {
	var sayi1, sayi2 int // Go dilinde değişken tanımlanırken -> var degisken_ad degisken_tip    şeklinde tanımlanır
	sayi1 = 3
	sayi2 = 5
	var toplam int  = sayi1 + sayi2 
	fmt.Println("Sayıların toplamı: ", toplam)


	var bolum float32 // sonucu float istediğimizden işleme sokacağımız sayıları da float olarak belirledik

	bolum = 8.0 / 32.4

	fmt.Println("Bölüm = ",bolum)

	var ad,soyad string

	ad = "Enes" // stringler ya çift tırnak
	soyad =  `Baysan` // ya da ters tırnak arasında tanımlanır

	fmt.Println(ad,"\t",soyad)

	/*
	\t  -> tab boşluk
	\n  -> alt satıra
	\"  -> çift tırnak işareti
	\`  -> tek tırnak işareti
	\\  -> ters bölü
	*/
}
