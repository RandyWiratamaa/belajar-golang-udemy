package main

import "fmt"

type Alamat struct {
	City, Province, Country string
}

func ChangeCountry(alamat *Alamat) {
	alamat.Country = "Indonesia"
}

func main() {
	var alamat = Alamat{
		City:     "Bandung",
		Province: "Jawa Barat",
		Country:  "Indonesia",
	}
	ChangeCountry(&alamat)
	fmt.Println(alamat)
}
