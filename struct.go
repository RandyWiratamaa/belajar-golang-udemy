package main

import "fmt"

type Customer struct {
	Name, Address, Gender string
	Age                   int
}

func (customer Customer) sayHello() {
	fmt.Println("Hello, My Name is", customer.Name)
	fmt.Println("i'm from", customer.Address)
	fmt.Println("i'm", customer.Age, "years old")
}

func main() {
	var biodata Customer
	biodata.Name = "Randy Wiratama"
	biodata.Address = "Jakarta"
	biodata.Gender = "Pria"
	biodata.Age = 27

	fmt.Println(biodata)

	//atau dengan cara Struct Literals
	bio := Customer{
		Name:    "Randy",
		Address: "Jakarta",
		Gender:  "Pria",
		Age:     27,
	}
	bio.sayHello()
}
