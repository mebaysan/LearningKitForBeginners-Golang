package main

import (
	"fmt"
)

func main() {

	var plaka int = 34
	switch plaka {
	case 34:
		fmt.Println("İstanbul")
	case 35:
		fmt.Println("İzmir")
	default: // yukarıdakilerden hiç biri olmazsa bu blok çalışır
		fmt.Println("Bilinmeyen plaka")
	}

}
