package structs

import "fmt"

func Demo1() {
	fmt.Println(product{"Bilgisayar", 12000, "HP"})
}

type product struct {
	name      string
	unitPrice float32
	brand     string
}
