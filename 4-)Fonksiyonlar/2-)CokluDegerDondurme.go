package main

import "fmt"

func main() {
	a := 1
	b := 2
	fmt.Printf("Değişimden Önce a -> %d\t Değişimden Önce b -> %d\n", a, b)
	a, b = degistir(a, b) // aşağıdan gelen b değişkeni burdaki a'ya eşitlendi, gelen a değişkeni burdaki b'ye eşitlendi
	fmt.Printf("Değişen a -> %d\t değişen b -> %d\n", a, b)
	// Fonksiyon kullanmadan da bu işlemi gerçekleştirebilirdik
	c := 3
	d := 4
	c, d = d, c // değişkenler yer değiştirdi
	fmt.Println(c, d)
}

func degistir(a, b int) (int, int) { // fonksiyonu tanımlarken 2 adet int tipinde değer return edeceğini belirttik -> (int,int)

	return b, a // önce b sonra a değişkenini return ettik

}
