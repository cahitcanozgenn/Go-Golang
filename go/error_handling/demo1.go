package error_handling

import (
	"fmt"
	"os"
)

func Demo1() {
	f, error := os.Open("demo1.txt")

	//dosya bulunursa f olur. error nil olur.
	// hata yoksa nil
	if error != nil {
		fmt.Println("Dosya Bulunamadı")
	} else {
		fmt.Println(f.Name())
	}

}
