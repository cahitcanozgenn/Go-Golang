package for_range

import "fmt"

func Demo1() {
	sehirler := []string{"Ankara", "Rize"}
	for _, sehir := range sehirler {
		fmt.Println(sehir)
	}
	// i kullanılmak istenmiyorsa _ kullanılabilir.
}
