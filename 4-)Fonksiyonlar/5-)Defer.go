package main

import "fmt"

/*
fonksiyon sonlanırken çalışmasını istediğimiz satırları 'defer' ile ifade ederiz
Eğer 1'den fazla 'defer' var ise LIFO mantığı ile çalışır. Yani son giren ilk çıkar. En sondaki 'defer' en önce çalışır, en üstteki 'defer' en son çalışır.
Step by step:
	-) Önce fonksiyon çalışacak ->  fmt.Println("Toplam = ", toplam)
	-) defer yolcula() çalışacak çünkü en sonuncu defer
	-) defer selamla() daha sonra çalışacak çünkü ilk eklenen defer
*/
func main() {
	defer selamla()
	toplam := topla(1, 2)
	fmt.Println("Toplam = ", toplam)
	defer yolcula()
}

func selamla() {
	fmt.Println("Selamlar  hoşgeldin...")
}

func yolcula() {
	fmt.Println("Bay bay iyi ki geldin...")
}

func topla(a, b int) int {
	cevap := a + b
	return cevap
}
