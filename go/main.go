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
	kdv3 := 53.55
	fmt.Println(kdv3)

	fmt.Printf("Veri Tipi: %T", kdv3) // değişkenin veri tipini öğrenme

}
