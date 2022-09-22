package interfaces

import "fmt"

type Person struct {
	firstName      string
	lastName       string
	identityNumber string
}

type Worker struct {
	firstName      string
	lastName       string
	identityNumber string
	workerNumber   string
}

type Customer struct {
	firstName      string
	lastName       string
	identityNumber string
	customerNumber string
}

func add(w Worker) {
	fmt.Println(w.firstName, "Adlı Personel Eklendi.")
}

func Demo3() {
	add(Worker{firstName: "Cahit Can", lastName: "ÖZGEN", identityNumber: "123456789", workerNumber: "12654"})
}
