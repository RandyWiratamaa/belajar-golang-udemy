package main

import "fmt"

/**
interface kosong tidak memiliki deklarasi method satupun.
secara otomatis semua tipe data akan menjadi implementasinya
eg : fmt.PrintLn, panic, recover
*/

//Check merupakan interface kosong
func Check(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "check 123"
	}
}

func main() {
	var data interface{} = Check(2)
	fmt.Println(data)
}
