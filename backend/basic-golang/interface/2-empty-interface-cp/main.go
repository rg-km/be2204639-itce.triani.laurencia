package main

import "fmt"

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan empty interface.
// Buatlah dua data makanan dan minuman yaitu ayam goreng dan cola yang memiliki atribut
// Nama, Jenis, dan Harga.
// Ayam Goreng, Cepat saji, 20000
// Cola, Minuman, 7000

func GetMenu() []map[string]interface{} {
	var menu []map[string]interface{}

	// TODO: answer here
<<<<<<< HEAD
<<<<<<< HEAD
	ayamGoreng := make(map[string]interface{})
	ayamGoreng["Nama"] = "Ayam Goreng"
	ayamGoreng["Jenis"] = "Cepat saji"
	ayamGoreng["Harga"] = 20000

	menu = append(menu, ayamGoreng)

	cola := make(map[string]interface{})
	cola["Nama"] = "Cola"
	cola["Jenis"] = "Minuman"
	cola["Harga"] = 7000

	menu = append(menu, cola)
=======
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
=======
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015

	return menu
}

func main() {
	menu := GetMenu()

	for _, m := range menu {
		for k, v := range m {
			fmt.Printf("%s: %v\n", k, v)
		}
	}
<<<<<<< HEAD
<<<<<<< HEAD
}
=======
}
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
=======
}
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
