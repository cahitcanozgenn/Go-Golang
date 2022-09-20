package main // dosyayı hangi klasöre yerleştireceğimizi belirliyoruz.
// fmt paketini kullanmak için projemize import ediyoruz.
//"golesson/conditionals"
//"golesson/loops"

import (
	"fmt"
	"golesson/channels"
	"golesson/functions"
)

func main() {
	//variables.Demo1()
	//conditionals.Demo2()
	//conditionals.Demo3()
	//loops.Demo2()
	//array.Demo4()

	//ctrl k ctrl c comment
	// sonuc1, sonuc2, sonuc3, sonuc4 := functions.DortIslem(10, 6)
	// fmt.Println("Toplam: ", sonuc1)
	// fmt.Println("Çıkarım: ", sonuc2)
	// fmt.Println("Çarpım: ", sonuc3)
	// fmt.Println("Bölüm: ", sonuc4)

	fmt.Println((functions.ToplaVariadic(1, 4, 6)))
	fmt.Println((functions.ToplaVariadic(1, 4, 6, 3, 6, 5)))

	// numbers := []int{4, 5, 69, 36}
	// fmt.Println((functions.ToplaVariadic(numbers...)))
	// sayi := 5
	// pointers.Demo2()
	// fmt.Println("Maindeki Sayı: ", sayi)

	// sayilar := []int{1, 2, 3}
	// pointers.Demo2(sayilar)
	// fmt.Println("Maindeki Sayı: ", sayilar[0])
	ciftSayiCn := make(chan int)
	tekSayiCn := make(chan int)
	go channels.CiftSayilar(ciftSayiCn)
	go channels.TekSayilar(tekSayiCn)
	// time.Sleep(5 * time.Second)
	ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn
	carpim := ciftSayiToplam * tekSayiToplam
	fmt.Println(carpim)
}
