package main

import "fmt"

/**
Sangat direkomendasikan menggunakan pointer di method (struct),
sehingga tidak boros memory karena harus selalu diduplikasi ketika memanggil method
*/

type Title struct {
	Name string
}

func (title *Title) Golang() {
	title.Name = "Golang Programming. " + title.Name
}

func main() {
	val := Title{"Pointer Method"}
	val.Golang()

	fmt.Println(val.Name)
}
