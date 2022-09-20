// asenkron programlama işlerin aynı anda yapılmasına imkan sağlamaktadır.
// Bir uygulama inerken arada işlerimize devam etmek.

package goroutines

import (
	"fmt"
	"time"
)

func CiftSayilar() {
	for i := 0; i <= 10; i += 2 {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func TekSayilar() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}
}
