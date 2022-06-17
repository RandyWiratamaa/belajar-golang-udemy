package helper

import "fmt"

/**
Variable atau Function yang diawali huruf besar bisa diakses dari luar package
Variable atau Function yang diawali huruf keci tidak bisa diakses dari luar package
*/
var Application = "Belajar Golang Udemy"
var version = 1 //tidak bisa diakses dari luar package
func Setting(name string) {
	fmt.Println("Bagian 46 :", name)
}
