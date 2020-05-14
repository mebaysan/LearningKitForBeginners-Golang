package main

import "fmt"

func main() {

	var flag, dogru_mu bool // default olarak flag -> false

	dogru_mu = true

	fmt.Println(flag || dogru_mu)
	fmt.Println("Baysan" == "baysan")
}

/*
&& -> VE

|| -> VEYA

!  -> DEĞİL

*/
