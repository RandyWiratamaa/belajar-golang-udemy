package main

import "fmt"

/**
Type Assertions adalah kemampuan merubah tipe data menjadi tipe data yang diinginkan
Sering sekali digunakan ketika bertemu dengan interface kosong
*/
func random() interface{} {
	return "Materi Type Assertions"
}

func main() {
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	//Penggunaan switch pada type assertions (Best Practice)
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown")
	}
}
