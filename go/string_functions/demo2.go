package string_functions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	isim := "Cahit"
	fmt.Println(s.HasPrefix(isim, "Cad")) // cad ile başlıyormu
	fmt.Println(s.HasSuffix(isim, "hit")) //  hit ile bitiyormu

}
