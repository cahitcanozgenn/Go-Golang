package conditionals

import "fmt"

func Demo1() {

	hesap := 1000
	cekilmekIstenen := 1900

	if cekilmekIstenen > hesap {
		fmt.Print("Hesaptaki Para Yetersiz")
	}
	//%f float %v value
	fmt.Println("Bitti. Hesaptaki Para: " + fmt.Sprint("%v", hesap))
	fmt.Printf("Hesaptaki Para: %v", hesap)
}
