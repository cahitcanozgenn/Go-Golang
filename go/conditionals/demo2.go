package conditionals

import "fmt"

func Demo2() {
	hesap := 1000
	cekilmekIstenen := 1900

	if cekilmekIstenen > hesap {
		fmt.Print("Hesaptaki Para Yetersiz")
	} else {
		fmt.Print("Paranız Hazırlanıyor")
	}

	/*
		if sart{
			if sart{

			}

		}
	*/
}
