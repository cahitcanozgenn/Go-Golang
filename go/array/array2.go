package array

import "fmt"

func Demo2() {
	var cities [5]string
	cities[0] = "Samsun"
	cities[1] = "Ankara"
	cities[2] = "Rize"
	cities[3] = "Antalya"
	cities[4] = "Adana"
	fmt.Println(cities)

	for i := 0; i < 5; i++ {
		fmt.Println(cities[i])
	}
}
