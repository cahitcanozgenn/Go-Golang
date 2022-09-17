package conditionals

import "fmt"

func Demo3() {
	// 3 adet int değişken tanımlayalım.
	// Ekrana en büyük olanı yazdırınız.

	number1 := 5
	number2 := 8
	number3 := 1
	enBuyuk := number1
	if number2 > enBuyuk {

		enBuyuk = number2
	}
	if number3 > enBuyuk {
		enBuyuk = number3
	}

	fmt.Printf("En Büyük Sayı: %v", enBuyuk)

}
