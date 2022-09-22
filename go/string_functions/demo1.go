package string_functions

import (
	"fmt"
	s "strings"
)

func Demo1() {

	isim := "Cahit"
	fmt.Println(s.Count(isim, "a"))
	fmt.Println(s.Contains(isim, "h"))
	fmt.Println(s.Index(isim, "h"))
}
