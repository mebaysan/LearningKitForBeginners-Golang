package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	contains()
	kac_kelime_var()
	parcala()
	basinda_sonunda_ara()
	nerede_basliyor()
	birlestir()
	haritaya_gore_degistir()
	tekrarla()
	degistir()
	ayraca_gore()
	buyuk_kucuk_cevir()
}

func contains() {
	mailler := []string{
		"mailadresi@gmail.com",
		"mail.com",
		"mailadresi2@gmail.com"}

	for _, mail := range mailler {
		if !strings.ContainsAny(mail, "@") { // 1. parametre hangi stringde aranacak, 2. parametre hangi ifade aranıyor
			fmt.Println("Hatalı mail adresi -> ", mail)
		}
	}
}

func kac_kelime_var() {
	// uzun stringleri ` ` arasına yazabiliriz
	cumle := `
	It is a long established fact that a reader will be distracted by the readable content 
	of a page when looking at its layout. The point of using Lorem Ipsum is that it has a 
	more-or-less normal distribution of letters, as opposed to using 'Content here, content here', 
	making it look like readable English. Many desktop publishing packages and web page editors now use 
	Lorem Ipsum as their default model text, and a search for 'lorem ipsum' will uncover many web sites 
	still in their infancy.`
	kelime_sayisi := strings.Count(cumle, " ") // 1. parametre hangi stringde aranacak, 2. parametre hangi string aranacak
	satir_sayisi := strings.Count(cumle, "\n")
	fmt.Printf("Cümledeki kelime sayısı -> %d\tsatır sayısı -> %d\n", kelime_sayisi, satir_sayisi)
}

func parcala() {
	cumle := `
	It is a long established fact that a reader will be distracted by the readable content 
	of a page when looking at its layout. The point of using Lorem Ipsum is that it has a 
	more-or-less normal distribution of letters, as opposed to using 'Content here, content here', 
	making it look like readable English. Many desktop publishing packages and web page editors now use 
	Lorem Ipsum as their default model text, and a search for 'lorem ipsum' will uncover many web sites 
	still in their infancy.`

	kelimeler := strings.Fields(cumle) // stringi whitespace denilen karakterlere göre böler -> \n \t ' ' gibi karakterler
	fmt.Printf("Toplamda %d kelime var\n", len(kelimeler))
	fmt.Printf("%q\n", kelimeler)
	// for _, kelime := range kelimeler {
	// 	fmt.Println(kelime)
	// }
}

func basinda_sonunda_ara() {
	kelime := "Ankara"
	fmt.Println(strings.HasPrefix(kelime, "a")) // false döndü, büyük-küçük harf duyarlılığı vardır, metin 1. parametre ile başlıyor mu
	fmt.Println(strings.HasSuffix(kelime, "a")) // true döner. metin 2. parametre ile bitiyor mu
}

func nerede_basliyor() {
	kelime := "Ali atabin. Enes ekmel almaya gitti mi? Yusuf akşam eve gelsin!"
	fmt.Println(strings.Index(kelime, "b"))        // string içerisinde aranan eleman ilk olarak nerde yakalanırsa index numarasını verir
	fmt.Println(strings.IndexAny(kelime, "?,!;:")) // string içerisinde aranan elemanLAR (2.parametredeki tüm elemanları tek tek döner) ilk olarak nerede kesişirse oranın index numarasını döner
}

func birlestir() {
	kelimeler := []string{"Ali", "ata", "bin", "!"}
	fmt.Println(strings.Join(kelimeler, ",")) // slice olarak verilen ifadeleri istenilen parametre ile birleştirir
	fmt.Println(strings.Join(kelimeler, "\t"))
}

func haritaya_gore_degistir() {
	rot13 := func(harf rune) rune { // rune veri tipi int32 (32 bitlik tamsayı)'nin diğer adıdır. ASCII karakter kodlama için 8 bit (0 - 255) yeterli olmaktadır. UTF-8 karakter kodlaması için 4 bayt(32 bit) gerekmektedir.
		return harf + 13
	}

	fmt.Println((strings.Map(rot13, "abc")))
}

func tekrarla() {
	for i := 0; i <= 5; i++ {
		fmt.Println(strings.Repeat("*", i)) // 1. parametredeki string'i 2. parametredeki kadar tekrarlar
	}
}

func degistir() {
	cumle := "kedileri çok severim, sokaktaki kedilere yiyecek veririm"
	fmt.Println(strings.Replace(cumle, "kedi", "köpek", -1)) // 4 parametre alır. 1 -> hangi string değişecek  2 -> stringdeki hangi ifadeler değişecek   3 -> neyle değişecek  4 -> kaç tanesi değişecek, eğer -1 verilirse hepsi değişir
	fmt.Println(strings.Replace(cumle, "kedi", "köpek", 1))
}

func ayraca_gore() {
	fmt.Printf("%q\n", strings.Split("a-b-c-d-e", "-"))      // klasik split, metni ayraca göle parçalar
	fmt.Printf("%q\n", strings.SplitAfter("a-b-c-d-e", "-")) // metni ayracın hemen sonrasından ayırır, ayraçları ifadelere dahil eder
	fmt.Printf("%q\n", strings.SplitN("a-b-c-d-e", "-", 3))  // metni N tane parçaya ayırır, N sayısı argüman olarak belirtilmelidir.
	fmt.Printf("%q\n", strings.SplitAfterN("a-b-c-d-e", "-", 3))
}

func buyuk_kucuk_cevir() {
	metin := "mErhaba enEs baYSan nasILsın?"
	fmt.Println(strings.Title(metin))   // kelimelerin ilk harfini büyültür
	fmt.Println(strings.ToLower(metin)) // bütün harfleri küçültür
	fmt.Println(strings.ToUpper(metin)) // küçük harfleri büyültür
	// türkçe karakterleri çevirirken sıkıntı çıkartıyor. Bunun için "unicode" paketi kullanılmaktadır
	fmt.Println("-------------------------")
	tr := unicode.TurkishCase
	fmt.Println(strings.ToTitleSpecial(tr, metin)) // ilk parametre hangi dile göre unicode yapacağımızdır
	fmt.Println(strings.ToLowerSpecial(tr, metin)) // ilk parametre hangi dile göre unicode yapacağımızdır
	fmt.Println(strings.ToUpperSpecial(tr, metin)) // ilk parametre hangi dile göre unicode yapacağımızdır
}
