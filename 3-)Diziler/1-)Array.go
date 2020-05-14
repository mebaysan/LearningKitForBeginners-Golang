package main

import "fmt"

func main() {
	// dizi içerisinde FARKLI TÜRDEN VERI TUTULAMAZ

	var iller [82]string // kaç elemandan oluşacağını ve tipini belirttik. Dizi tanımlandığı anda tüm elemanları için kullanılmasa bile bellekte yer ayrılmıştır.
	iller[1] = "Adana"   // Go'da da diziler 0. indisten başlar
	iller[34] = "Adana"
	fmt.Println(iller)
	fmt.Println(iller[34]) // dizi içerisindeki elemanlara indis değeri kullanılarak erişilebilir

	var adaylar [5]float32
	adaylar[0] = 15
	adaylar[1] = 24
	adaylar[2] = 18
	adaylar[3] = 67
	adaylar[4] = 76

	var toplam float32

	for i := 0; i < len(adaylar); i++ {
		toplam += adaylar[i] // python'daki gibi += yapabiliriz, len(adaylar) diyerek listenin eleman sayısını bulabiliriz
	}
	fmt.Println("Toplam aday sayısı -> ", len(adaylar))
	fmt.Println("Adayların puan ortalaması -> ", toplam/5)

	notlar := [5]int{56, 34, 12, 67, 98}

	for z := 0; z < len(notlar); z++ {
		if notlar[z] < 50 {
			fmt.Println("Aday başarısız oldu -> ", notlar[z])
		}
	}

	puanlar := [5]int{56, 34, 12, 67, 98}

	for i, puan := range puanlar { // range çok sayıda veri üzerinde for'u döndürmek için kullanılır
		if puan < 50 { // bu kullanımda i yine döngü değişkenidir ve her kullanımda artmaktadır, puan ise o an döndüğü elemandır (puanlar[i])
			fmt.Printf("%d. aday başarısız oldu\n", i)
		}
	}

	puanlar = [5]int{56, 34, 12, 67, 98}
	basarili := 0
	for _, puan := range puanlar { // eğer sayaç değişkeninin bizim için bir önemi yoksa '_' olarak kullanabiliriz. Yukarda i. aday diyorduk burada hiç kullanmadık
		if puan >= 50 {
			// basarili++
			basarili += 1
		}
	}
	fmt.Printf("%d öğrenci başarılı\n", basarili)
	fmt.Println("=======================================")

	// MATRISLER

	matris := [3][2]int{{11, 12}, {21, 15}, {16, 5}}

	fmt.Println("Matrisin tamamı -> ", matris)
	for satir := 0; satir < len(matris); satir++ {
		for sutun := 0; sutun < 2; sutun++ {
			fmt.Print(matris[satir][sutun], " ")
		}
		fmt.Println()
	}

}
