package main

import "fmt"

// Go'da kalıtım'ın karşılığı 'embedded types'dır.
// Yapı şeklinde tanımlanan (struct) bir türe ait elemanları (fields & methods) başka bir yapı için de kullanabilmeyi sağlar
func main() {

	hayvan := Hayvan{isim: "Köpek", yas: 4, ayak: 4}
	var muhabbet Kus
	muhabbet.isim = "Muhabbet Kuşu"
	muhabbet.yas = 2
	muhabbet.ayak = 2
	muhabbet.kanat = 2 
	hayvan.nefes_al()
	muhabbet.Hayvan.nefes_al() // 1'den fazla embedded types olsaydı ve hepsinsen 'nefes_al' metotu gelseydi, bunun nereden geldiğini belirtmemiz gerekirdi -> muhabbet.HAYVAN.nefes_al(), anyı şey field'lar için de geçerli
	muhabbet.nefes_al() // tek bir yerden nefes_al geldiği için belirtmesek de olur
}

type Hayvan struct { // Hayvan adında bir yapı tanımladık
	isim      string
	yas, ayak int
}

type Kus struct {
	Hayvan // Kus yapısı Hayvan yapısından türetilmiştir dedik (embedded types). İstenildiği kadar embedded types eklenebilir
	kanat  int
}

func (self Hayvan) nefes_al() {
	fmt.Printf("%s nefes alıyor...\n", self.isim)
}
