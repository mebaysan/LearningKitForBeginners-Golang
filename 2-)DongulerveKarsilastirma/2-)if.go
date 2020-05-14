package main

import (
	"fmt"
)

func main() {
	var i, b2, b3, be, b4 int

	for i = 1; i <= 100; i++ {
		if i%2 == 0 { // 2'ye bölünebiliyorsa
			b2++
			if i%4 == 0 { // 4'e bölünenler
				b4++
			}
		} else if i%3 == 0 { // 3'e bölünebiliyorsa
			b3++
		} else { // yukarıdakilerden hiç biri değilse
			be++
		}

	}

	fmt.Printf("2:%d - 3:%d - diğerleri:%d - 4:%d", b2, b3, be, b4) // printf string formatlamak için kullanılır. %d oraya sayısal değer(digit) yazılacağını ifade eder
	fmt.Println("\n")

}
