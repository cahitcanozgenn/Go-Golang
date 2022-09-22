package interfaces

import "fmt"

// Konut Kredisi
type Mortgage struct {
	creditPaymentTotal float64
	address            string
	rate               float64
}

// Taşıt Kredisi
type Car struct {
	creditPaymentTotal float64
	carInfo            string
	rate               float64
}

type CreditCalculator interface {
	Calculate() float64
}

func (m Mortgage) Calculate() float64 {
	return m.creditPaymentTotal * m.rate / 12
}

func (c Car) Calculate() float64 {
	return c.creditPaymentTotal * c.rate / 12
}

func CalculateMonthlyPayment(credits []CreditCalculator) float64 {
	paymentTotal := 0.0
	for i := 0; i < len(credits); i++ {
		paymentTotal = paymentTotal + credits[i].Calculate()
	}
	return paymentTotal

}

func Demo2() {
	credit1 := Mortgage{rate: 10, creditPaymentTotal: 100000, address: "SAMSUN/ATAKUM"}
	credit2 := Mortgage{rate: 12, creditPaymentTotal: 500000, address: "İSTANBUL"}
	credit3 := Car{rate: 20, creditPaymentTotal: 5000000, carInfo: "Mercedes-Benz Travego 17-SHD"}

	credits := []CreditCalculator{credit1, credit2, credit3}
	total := CalculateMonthlyPayment(credits)
	fmt.Println("Toplam Ödeme: ", total)
}
