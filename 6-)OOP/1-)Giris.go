package main

import "fmt"

// Go dilinde; bilinen diğer dillere nazaran 'class' ön eki 'struct' ile yer değiştirmiştir.
// Person class'ından demek yerine Person yapısı'ndan diyebiliriz.

func main() {
	insan1 := Insan{"Enes", "Baysan", 20} // instance oluşturduğumuz anda field'larının değerlerini sırasıyla set edebiliriz
	insan2 := Insan{soyad: "Baysan", ad: "Yusuf", yas: 19} // istersek key:value şeklinde set edebiliriz
	var insan3 Insan // Insan tipinde insan3 değişkeni, istersek de önce instance oluşturur sonra field'larını set ederiz
	insan3.ad = "Yavuz"
	insan3.soyad = "Baysan"
	insan3.yas = 41
	fmt.Printf("3 insanın yaşları toplamı -> %d\n", yaslari_topla(insan1.yas, insan2.yas, insan3.yas))
}

type Insan struct { // bir veri tipi oluşturduk
	ad, soyad string
	yas       int
}

func yaslari_topla(yaslar ...int) int {
	toplam := 0
	for _, yas := range yaslar {
		toplam += yas
	}
	return toplam
}
