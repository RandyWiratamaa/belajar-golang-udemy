package main

import "fmt"

/**
kontrak baru dengan bentuk interface
interface dengan nama HasName
*/
type HasName interface {
	//mempunyai method GetName dengan return string
	GetName() string
}

/**
interface dapat digunakan sebagai parameter pada function
func sayHello memiliki parameter dari interface (HasName)
*/
func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

//Struct
type Person struct {
	Name string
}

//Struct yang berbeda bisa menggunakan interface yang sama
type Animal struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var nama Person
	nama.Name = "Randy"

	sayHello(nama)

	cat := Animal{
		Name: "Push",
	}

	sayHello(cat)
}
