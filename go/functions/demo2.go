package functions

func DortIslem(number1 int, number2 int) (int, int, int, float32) {
	toplam := number1 + number2
	cikarim := number1 - number2
	carpim := number1 * number2
	bolum := number1 / number2
	return toplam, cikarim, carpim, float32(bolum)
}
