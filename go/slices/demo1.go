package slices

import "fmt"

func Demo1() {
	names := make([]string, 3)

	names[0] = "Cahit"
	names[1] = "Can"
	names[2] = "ÖZGEN"
	names = append(names, "Developer") // diziye yeni eleman ekleme. Append=eklemek
	fmt.Println(names)
}
