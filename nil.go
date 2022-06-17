package main

import "fmt"

//nil hanya bisa digunakan pada inteface, function, map, slice, pointer dan channel
func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("Materi Nil")
	fmt.Println(data)
}
