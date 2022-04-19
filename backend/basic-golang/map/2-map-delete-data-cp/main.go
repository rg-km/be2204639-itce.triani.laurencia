package main

import "fmt"

// pada cekpoin ini kalian akan mencoba untuk menghapus data dengan tipe []map[string]string
// gabungan slice dan map.

func main() {
	var namaUmur = []map[string]string{
		{"name": "Socrates", "gender": "male"},
		{"name": "Plato", "gender": "male"},
		{"name": "Ada", "gender": "female"},
		{"name": "Leonhard Euler", "gender": "female"},
		{"name": "Blaise Pascal", "gender": "male"},
	}

	// terdapat kesalahan pada data gender tersebut dapatkan kalian memperbaiki nya ?
	// TODO: answer here
	for _, val := range namaUmur {
<<<<<<< HEAD
<<<<<<< HEAD
		fmt.Println(val["name"], " ", val["gender"],)	
	}
	fmt.Println()
=======
		fmt.Println(val["name"], " ", val["gender"])
	}

>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
=======
		fmt.Println(val["name"], " ", val["gender"])
	}

>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
	// Nah coba saatnya kalian menghapuskan key "gender" pada setiap data
	// delete data if key is equal "gender"

	for _, val := range namaUmur {
		fmt.Println(val)
	}
<<<<<<< HEAD
<<<<<<< HEAD
	fmt.Println()
=======
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
=======
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
	// Output sebelum dihapus
	/*
		map[gender:male name:Socrates]
		map[gender:male name:Plato]
		map[gender:female name:Ada]
		map[gender:male name:Leonhard Euler]
		map[gender:male name:Blaise Pascal]
	*/

	// TODO: answer here
<<<<<<< HEAD
<<<<<<< HEAD
	for _, val := range namaUmur {
		delete(val, "gender")
		fmt.Println(val)
	}
	fmt.Println()
=======
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
=======
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015

	// Output setelah dihapus
	/*
		map[name:Socrates]
		map[name:Plato]
		map[name:Ada]
		map[name:Leonhard Euler]
		map[name:Blaise Pascal]
	*/
<<<<<<< HEAD
<<<<<<< HEAD
	
=======
	for _, val := range namaUmur {
		fmt.Println(val)
	}
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
=======
	for _, val := range namaUmur {
		fmt.Println(val)
	}
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
}
