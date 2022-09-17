package loops

import "fmt"

func Demo1() {
	i := 1
	for i < 6 {

		fmt.Println("Cahit")
		i = i + 1
	}
	fmt.Println("Bitti")
}

func Demo2() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
