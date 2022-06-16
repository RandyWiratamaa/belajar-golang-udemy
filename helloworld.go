package main

import "fmt"

func main() {
	//String
	fmt.Println("Go Lang Udemy")

	//Function Return Values
	a, b := returnValues()
	fmt.Println(a, b)

	lenString()
	number()
	boolean()
	deklarasiVariable()
	constant()
	konversiTipeData()
	deklarasiTipeData()
	tipeDataArray()
	tipeDataSlice()
	tipeDataMap()
	forLoops()

	data_slice := []int{10, 20, 30, 40, 50}
	total := variadicFunction(data_slice...)
	fmt.Println(total)

}

func lenString() {
	fmt.Println("---------------------------")
	fmt.Println("Menghitung Panjang Karakter")
	fmt.Println("---------------------------")
	fmt.Println(len("Hello World"))
	fmt.Println("=====================================")
}

func number() {
	fmt.Println("------")
	fmt.Println("Number")
	fmt.Println("------")
	fmt.Println("Satu =", 1)
	fmt.Println("Dua =", 2)
	fmt.Println("Tiga koma lima =", 3.5)
	fmt.Println("=====================================")
}

func boolean() {
	fmt.Println("-------")
	fmt.Println("Boolean")
	fmt.Println("-------")
	fmt.Println("Benar =", true)
	fmt.Println("false =", false)
	fmt.Println("=====================================")
}

func deklarasiVariable() {
	fmt.Println("------------------")
	fmt.Println("Deklarasi Variable")
	fmt.Println("------------------")
	name := "Randy Wiratama"
	fmt.Println(name)
	fmt.Println("=====================================")
}

func constant() {
	fmt.Println("----------------------------")
	fmt.Println("Constant					 ")
	fmt.Println("Nilai yang tidak bisa diubah")
	fmt.Println("----------------------------")
	//Langsung di deklarasikan datanya
	var (
		firstName = "Randy"
		lastName  = "Wiratama"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)

	const (
		nama = "Randy"
		umur = 27
	)
	fmt.Println(nama)
	fmt.Println(umur)
	fmt.Println("=====================================")
}

func konversiTipeData() {
	fmt.Println("-------------------")
	fmt.Println("Konversi tipe data ")
	fmt.Println("-------------------")
	// var [nama_variable] [tipe_data] = [tipe_data]([nama_variable])
	var nilai001 int32 = 10000
	var nilai_convert int64 = int64(nilai001)
	fmt.Println(nilai_convert)
	fmt.Println("=====================================")
}

func deklarasiTipeData() {
	fmt.Println("----------------")
	fmt.Println("Type Declaration")
	fmt.Println("----------------")
	// type [alias_tipeData] [tipe_data]
	type noKTP string
	var KTPRandy noKTP = "123456789"
	fmt.Println(KTPRandy)
	fmt.Println("=====================================")
}

func tipeDataArray() {
	fmt.Println("---------------")
	fmt.Println("Tipe Data Array")
	fmt.Println("---------------")
	var values = [...]int{
		10, 20, 30, 40, 50,
	}
	fmt.Println(values)
	fmt.Println("=====================================")
}
func tipeDataSlice() {
	fmt.Println("---------------")
	fmt.Println("Tipe Data Slice")
	fmt.Println("---------------")
	bulan := [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	slice := bulan[4:7]

	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])

	fmt.Println("----------------")
	fmt.Println("## Membuat Slice")
	fmt.Println("----------------")
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Go"
	newSlice[1] = "Lang"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	fmt.Println("-------------")
	fmt.Println("## Copy Slice")
	fmt.Println("-------------")
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
	fmt.Println("=====================================")
}
func tipeDataMap() {
	fmt.Println("--------------")
	fmt.Println("## Membuat Map")
	fmt.Println("--------------")

	biodata := map[string]string{
		"name":    "Randy",
		"address": "Jakarta",
		"gender":  "Laki-laki",
	}
	fmt.Println(biodata)
	fmt.Println(biodata["name"])
	fmt.Println(biodata["address"])
	fmt.Println(biodata["gender"])
	delete(biodata, "gender")
	fmt.Println(biodata)
	fmt.Println("=====================================")
}
func forLoops() {
	fmt.Println("-------------------------")
	fmt.Println("## Perulangan - For Loops")
	fmt.Println("-------------------------")

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	hari := []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	for i := 0; i < len(hari); i++ {
		fmt.Println(hari[i])
	}

	for i, value := range hari {
		fmt.Println("index", i, "=", value)
	}
}

func returnValues() (firstName, lastName string) {
	fmt.Println("-------------")
	fmt.Println("Return Values")
	fmt.Println("-------------")
	firstName = "Randy"
	lastName = "Wiratama"
	return
}

func variadicFunction(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}
