package main

import "fmt"

func main() {
	// sabitlerin değerleri değiştirilemez

	const ulke string = "TR" // sabitler tanımlanırken inline (tek satırda) olarak tanımlanır
	// değişken tanımlanırken aynı anda değeri de verilirse bu değere başlangıç değeri denir
	fmt.Println(ulke)
}
