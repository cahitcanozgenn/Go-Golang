package interfaces

import "fmt"

func dogrula(i interface{}) {
	sayi, ok := i.(int)
	fmt.Println(sayi, ok)
}

func Demo4() {
	dogrula(5)
}
