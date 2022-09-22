package error_handling

import (
	"fmt"
	"os"
)

// tyoe assertion
func Demo1() {
	f, error := os.Open("demoe1.txt")

	//dosya bulunursa f olur. error nil olur.
	// hata yoksa nil
	if error != nil {
		if pErr, ok := error.(*os.PathError); ok {
			fmt.Println("Dosya Bulunamadı", pErr)
			return
		} else {
			fmt.Println("Dosya Bulunamadı")
		}
		fmt.Println("Dosya Bulunamadı")
	} else {
		fmt.Println(f.Name())
	}

}
