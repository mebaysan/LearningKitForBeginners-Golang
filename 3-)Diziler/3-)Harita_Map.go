package main

import (
	"encoding/json"
	"fmt"
)

// key value şeklinde değerler tutulur
func main() {

	// var harita_ismi map[anahtar_veri_tipi]deger_veri_tipi

	// var il_plaka map[string]int // il_plaka adında map tanımladık
	// il_plaka = make(map[string]int) // il_plaka adında map oluşturduk

	il_plaka := make(map[string]int)
	il_plaka["İstanbul"] = 34 // pythondaki gibi eleman ekledik
	il_plaka["Ankara"] = 06
	il_plaka["İzmir"] = 35
	il_plaka["Adana"] = 01
	fmt.Println(il_plaka)
	fmt.Println(il_plaka["İstanbul"]) // bir map içerisinden elemana [key] şeklinde erişebiliriz

	json_tip, _ := json.MarshalIndent(il_plaka, "", " ") // verilerimizi jsoN'a çevirebiliriz
	fmt.Println(string(json_tip))

	for il, plaka:= range il_plaka{
		fmt.Printf("%s -> %d\n",il,plaka)
	}
	delete(il_plaka,"Ankara") // ilgili key'i ilgili mapten siler
	
}
