package main

import "fmt"

func main() {
	enes := Insan{boy: 182}
	fmt.Println("Enes'in ideal kilosu (metot) -> ", enes.idealkilo())
	fmt.Println("Enes'in ideal kilosu (fonksiyon) -> ", idealkilo(enes.boy))
	var x SAYI
	x = 3
	fmt.Println(x.kare_al())
}

type Insan struct {
	boy float32
}

// Insan yapı'sına özel bir metot tanımlamış olduk.
// Metot ve Fonksiyonun şekil olarak birbirinden farkı 'func sözcüğü ile metotun ismi arasındaki parantez içerisindeki kısımdır'. Bu kısım fonksiyonu (metotu) yapı'ya (struct) bağlar -> [[ (birisi Insan) ]]
func (self Insan) idealkilo() float32 { // metot
	kilo := 0.9*self.boy - 85
	return kilo
}

func idealkilo(boy float32) float32 { // fonksiyon
	kilo := 0.9*boy - 85
	return kilo
}

// Metotlar sadece yapılara özel değildir! Dilersek herhangi bir türe özel metot da oluşturabiliriz.

type SAYI int // int tipinde SAYI türünde. SAYI türü int türü ile aynı özelliklere sahip olacak

func (self SAYI) kare_al() int { // SAYI türü'ne bağlı bir metot
	return int(self * self)
}
