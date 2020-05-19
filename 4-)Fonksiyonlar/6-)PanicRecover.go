package main

import "fmt"

/*
'panic' ve 'recover' fonksiyonları go'da hazır olarak gelir. Bir fonksiyonda 'panic' oluşursa fonksiyonda 'defer' kısmı işletilip fonksiyondan çıkılır.
Hata durumunda kontrolü elde tutabilmek için 'panic' ve 'recover' ifadelerini kullanırız.
'recover' ise 'panic' durumunda oluşan hata durumunu yönetmemizi sağlar
*/

func main() {
	kare_al(5)
}

func kare_al(x int){
	defer func(){
		hata := recover()
		if hata!= nil { // hata mesajı var mı, yani panic() kullanılmış mı
			fmt.Println(hata)
			fmt.Println("Panik oluştu ve defer kısmı çalıştı...")
		}
	}()
	panic("durma noktası!")
	fmt.Println(x*x)
}