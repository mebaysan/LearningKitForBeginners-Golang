package main

import (
	"fmt"
	"os"
)

func main() {
	bilgileri_goster()
	dosya_icerik_oku()
}

func bilgileri_goster() {
	//dosya, _ := os.Open("1-)Strings.go") // açılan dosya; 'dosya' isimli değişkene aktarılır. '_' ise dönecek olan 2. değeri önemsemediğimizi ifade eder
	dosya, hata := os.Open("1-)Strings.go") // açılan dosya; 'dosya' isimli değişkene aktarılır. 2. parametre eğer varsa hata değişkenidir
	if hata != nil {
		fmt.Println("Hata var! -> ", hata)
		return
	}
	defer dosya.Close() // çıkarken dosyayı kapatalım
	dosya_bilgileri, _ := dosya.Stat()
	fmt.Println("Dosya ismi: ", dosya_bilgileri.Name())
	fmt.Println("Klasör mü: ", dosya_bilgileri.IsDir())
	fmt.Println("Yetkileri: ", dosya_bilgileri.Mode())
	fmt.Println("Değişim Tarihi: ", dosya_bilgileri.ModTime())
	fmt.Println("Boyut: ", dosya_bilgileri.Size(), "bayt")
}

func dosya_icerik_oku() {
	dosya, hata := os.Open("1-)Strings.go")
	if hata != nil {
		fmt.Println("Hata var! -> ", hata)
		return
	}
	defer dosya.Close()
	dosya_bilgisi, _ := dosya.Stat()
	byte_kesit := make([]byte, dosya_bilgisi.Size())
	okunan_byte, _ := dosya.Read(byte_kesit)
	okunan_metin := string(okunan_byte)
	fmt.Println(dosya.Name(), " isimli dosyadan okunan ", okunan_byte, " byte şu şekildedir:\n", okunan_metin)
}
