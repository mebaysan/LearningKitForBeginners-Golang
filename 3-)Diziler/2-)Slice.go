package main

import "fmt"

func main() {
	// eleman yapısı dinamik (değiştirilebilir) olan arraylere SLICE (kesit) denir
	// var iller []string                // string türünde boş bir slice (kesit) tanımladık
	// plakalar := make([]int, 82)       // 82 elemanlı bir slice (kesit) oluşturduk
	// cocuklar := make([]string, 2, 10) // 2 string içeren, 10'a kadar genişleyebilen slice (kesit) oluşturduk
	// // make -> kesit oluşturmak için kullanılır
	// // slice'lar bir diziye bağlı olduğuna göre, diziler üzerinden slice alabiliriz

	// var hayat []int
	// cocukluk := hayat[:15]
	// genclik := hayat[16:30]
	// ortayas := hayat[31:45]
	// ileriyas := hayat[45]
	// cap(array) -> kapasitesini verir
	// len(array) -> eleman sayısını verir
	var dizi = [4]string{"a", "b", "c", "d"}
	var kesit = dizi[:3]
	fmt.Printf("1-) eleman sayısı -> %d\tkapasite -> %d\t%s\n", len(kesit), cap(kesit), kesit)
	kesit = append(kesit, "x") // kesite bir eleman daha ekledi, kapasiteyi(4) aşmadı
	fmt.Printf("2-) eleman sayısı -> %d\tkapasite -> %d\t%s\n", len(kesit), cap(kesit), kesit)
	kesit = append(kesit, "y") // kesite bir eleman daha ekledi, kapasiteyi(4) aştı. Bu sebeple kapasite öncekinin 2 katına çıkarıldı. Bellekte yeni bir dizi oluşturuldu
	fmt.Printf("3-) eleman sayısı -> %d\tkapasite -> %d\t%s\n", len(kesit), cap(kesit), kesit)
	kesit = kesit[2:] // eleman sayısı 2 azaltıldı. Bu sebeple bellekte fazla yer kaplamaması için otomatik olarak array kapasitesi 8'den düşürülüp 6'ya çıktı, yeni değerin 2 katı
	fmt.Printf("4-) eleman sayısı -> %d\tkapasite -> %d\t%s\n", len(kesit), cap(kesit), kesit)

	ilk := []int{1, 2, 3}
	son := []int{4, 5}
	hepsi := append(ilk, son...) // son... -> son isimli slice'ın tüm elemanlarını,değişkenlerini ilk isimli slice'a ekle.
	fmt.Println(hepsi)

	kaynak := []int{1, 2, 3, 4}
	hedef := make([]int, 2)
	sayi := copy(hedef, kaynak) // kaç adet eleman kopyalandığını döner. sağdaki kesit soldaki kesite kopyalanır
	fmt.Printf("%d adet eleman kopyalandı\n", sayi)

}
