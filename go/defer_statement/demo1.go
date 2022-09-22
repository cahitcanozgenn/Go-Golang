package defer_statement

import "fmt"

// fonksiyonun sonunda çalışacak isimler
func A() {
	fmt.Println("A FONKSİYONU ÇALIŞTI")
}

func B() {
	defer A()
	defer C() // b nin bitiminden sonra ilk bu fonksiyon çalışacak.
	fmt.Println("B FONKSİYONU ÇALIŞTI")

}

func C() {
	fmt.Println("C FONKSİYONU ÇALIŞTI")
}
