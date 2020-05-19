package main

import "fmt"

// kendi kendini çağıran fonksiyonlardır
func main() {

	fmt.Println("Faktoriyel 5! -> ", faktoriyel(5))
}

func faktoriyel(sayi int) int {
	if sayi == 0 {
		return 1
	}
	return sayi * faktoriyel(sayi-1)
}
