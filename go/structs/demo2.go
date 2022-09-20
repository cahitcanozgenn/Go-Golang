package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

func (c customer) save() {
	fmt.Println(c.firstName, "Eklendi")
}

func (c customer) update() {
	fmt.Println(c.firstName, "Güncellendi")
}

func (c customer) delete() {
	fmt.Println(c.firstName, "Silindi")
}

func Demo2() {
	c := customer{firstName: "Cahit Can", lastName: "ÖZGEN", age: 20}
	c.save()
	c.delete()
	c.update()
}
