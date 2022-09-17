package array

import "fmt"

func Demo3() {
	sayilar := [5]int{20, 30, 40, 50, 20}
	fmt.Println(sayilar)

	//length=uzunluk
	for i := 0; i < len(sayilar); i++ {
		fmt.Println(sayilar[i])
	}
}
