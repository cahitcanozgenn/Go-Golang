package slices

import "fmt"

func Demo2() {
	cities := []string{"Ankara", "İstanbul", "İzmir"}
	fmt.Println(cities)
	citiesCopy := make([]string, len(cities))

	copy(citiesCopy, cities) // dizinin değerlerini kopyalama. Neyi nereye kopyalayayım.
	fmt.Println(citiesCopy)
	fmt.Println(cities[1:3])
}
