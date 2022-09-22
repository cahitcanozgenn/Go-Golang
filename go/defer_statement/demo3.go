package defer_statement

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) save() {
	fmt.Println("Ürün Kaydedildi", p.productName)

}

func Demo3() {
	p := product{productName: "Laptop", unitPrice: 1200}
	defer p.save()
	fmt.Println("İşlem Başarılı")

}
