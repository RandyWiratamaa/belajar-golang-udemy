package database

import "fmt"

/**
untuk membuat function yang dapat diakses secara otomatis ketika package diakses,
cukup membuat function dengan nama init
*/

var connection string

func init() {
	fmt.Println("Init berhasil otomatis diakses")
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
