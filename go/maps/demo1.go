package maps

import "fmt"

func Demo1() {
	// key-value
	sozluk := make(map[string]string)
	sozluk["table"] = "Masa"
	sozluk["book"] = "Kitap"

	fmt.Println(sozluk["book"])
	fmt.Println("Eleman Sayısı: ", len(sozluk))
	fmt.Println("Sözlük: ", sozluk)
	delete(sozluk, "book")
	fmt.Println("Eleman Sayısı: ", len(sozluk))
	fmt.Println("Sözlük: ", sozluk)

	value, varMi := sozluk["lamba"]
	fmt.Println(value)
	fmt.Println("Listede olma durumu: ", varMi)

	sozluk2 := map[string]string{"glass": "bardak", "microphone": "mikrofon"}
	fmt.Println(sozluk2)

}
