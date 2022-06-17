package main

import "fmt"

/**
secara default di Go-Lang semua variable di passing by value
jika mengirim sebuah variable ke dalam function, method atau variable lain,
sebenarnya yang dikirim adalah duplikasi valuenya (pass by value)

Pointer adl kemampuan membuat refernce ke lokasi data di memory yg sama,
tanpa menduplikasi data yg sudah ada.
Dengan pointer bisa membuat pass by reference

Untuk membuat sebuah variable dgn nilai pointer ke variable lain,
bisa menggunakan operator "&" (AND) diikuti dengan nama variablenya
*/

type Address struct {
	City, Province, Country string
}

func main() {
	//Pass By Value
	address1 := Address{
		City:     "Jakarta Selatan",
		Province: "Jakarta",
		Country:  "Indonesia",
	}

	address2 := address1
	address2.City = "Jakarta Barat"
	// fmt.Println(address1)
	// fmt.Println(address2)

	//Pass by reference (Pointer)
	address3 := &address1
	address3.City = "Jakarta Pusat"

	*address3 = Address{
		City:     "Malang",
		Province: "Jawa Timur",
		Country:  "Indonesia",
	}
	fmt.Println(address1)
	fmt.Println(address3)

	address4 := new(Address)
	address4.Province = "Sumatera Barat"
	fmt.Println(address4)
}
