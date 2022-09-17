package main // dosyayı hangi klasöre yerleştireceğimizi belirliyoruz.
// fmt paketini kullanmak için projemize import ediyoruz.
//"golesson/conditionals"
//"golesson/loops"
import (
	"fmt"
	"golesson/functions"
)

//"golesson/variables"

func main() {
	//variables.Demo1()
	//conditionals.Demo2()
	//conditionals.Demo3()
	//loops.Demo2()
	//array.Demo4()

	sonuc1, sonuc2, sonuc3, sonuc4 := functions.DortIslem(10, 6)
	fmt.Println("Toplam: ", sonuc1)
	fmt.Println("Çıkarım: ", sonuc2)
	fmt.Println("Çarpım: ", sonuc3)
	fmt.Println("Bölüm: ", sonuc4)
}
