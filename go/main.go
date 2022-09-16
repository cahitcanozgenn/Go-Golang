package main // dosyayı hangi klasöre yerleştireceğimizi belirliyoruz.
import "fmt" // fmt paketini kullanmak için projemize import ediyoruz.

func main() {

	// main bloğu go kodlarının çalışacağı bloktur.
	// go satır base'li bir dildir.
	// fmt=format
	// go büyük küçük harf duyarlı bir dildir.
	// kullanmadığın bir şeyi import edemezsin.
	// go kullanılmayan bir şeyi sevmez.
	var metin string = "Software Developer"
	fmt.Println(metin)
	fmt.Println("Cahit")

	var kdv int = 10
	fmt.Println(kdv)
	fmt.Println(100 + 100*10/100)
	fmt.Println()

	var kdv2 float32 = 0.1
	fmt.Println(kdv2)

	//değişken tanımlamanın bir başka yolu
	kdv3 := 53.55 // go dilinde en çok kullanılan değişken tanımlama tipidir.
	fmt.Println(kdv3)

	fmt.Printf("Veri Tipi: %T\n", kdv3) // değişkenin veri tipini öğrenme \n=yeni bir satır.

	var durum bool = true

	var metin1 string = "Arif"
	var metin2 string = "Cahit"

	// = değer atama işlemi yapar
	// == ?
	durum = metin1 == metin2 // metin1 metin2 ye eşit mi

	fmt.Println(durum)

}
