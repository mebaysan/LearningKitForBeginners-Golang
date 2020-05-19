package main

import "fmt"

func main() {
	var isim string
	fmt.Println("Adınız Nedir?")
	fmt.Scanln(&isim)
	fmt.Printf("Merhaba %s\n", isim)
}
