package main

import ( 
	"fmt" 
	"reflect" 
) 

const deneme = "DENEME" // bu bir global değişkendir
const deneme2 string = "DENEME2" // bu bir global değişkendir (2 tipte de sabit tanımlayabiliriz)

const (
	sabit1 int = 15
	sabit2 int = 25

) // bu şekilde çok sayıda değişken de tanımlayabiliriz

var degisen string = "Değişebilir" // sadece sabit tanımlamak zorunda değiliz
var degisen2 = "Değişebilir2" // sadece sabit tanımlamak zorunda değiliz

func main() {

	// aşağıdaki değişkenlerin her biri sadece bu fonksiyon için geçerli olacaklardır.
	
	var a int = 5 // 1. yöntem
	var b = 3 // 2. yöntem, burada değişkenin tipini sonradan algılaması için derleyiciye bıraktık
	c:= "Naber" // 3. yöntem, var ön eki de kaldırılarak kullanılabilir
	var d = "Go"
	const sabit int = 34

	fmt.Println("Toplam: ",a+b)
	fmt.Println(reflect.TypeOf(b),reflect.TypeOf(d))
	fmt.Println("c = ",c)
	c = "Naber Değişti" // sonraki tüm atamalarda yine '=' kullanılmalıdır
	fmt.Println("c = ",c)
	fmt.Println("Bu bir global değişkendir = ", deneme)
	fmt.Println("Bunlar global değişken toplamlarıdır = ", sabit1+sabit2)
	fmt.Println("Sadece sabit tanımlamak zorunda değiliz = ",degisen)
	fmt.Println("Sadece sabit tanımlamak zorunda değiliz = ",degisen2)

}

