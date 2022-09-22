package defer_statement

import "fmt"

func Demo2(sayi int32) string {
	if sayi%2 == 0 {
		return "Çift Sayıdır."
	}

	if sayi%2 != 0 {
		return "Tek Sayıdır."
	}
	return "Belli Değil"
}

func Test() {
	sonuc := Demo2(0)
	fmt.Println(sonuc)
}
