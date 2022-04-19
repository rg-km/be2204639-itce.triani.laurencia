package main

import "fmt"

// Untuk membaca slice itu sebenarnya sama seperti pada array.
// Kita cukup mengambil index nya
func main() {
<<<<<<< HEAD
<<<<<<< HEAD
	var slice []string
	slice = append(slice, "k")
	slice = append(slice, "a")
	slice = append(slice, "k")
	slice = append(slice, "a")
	slice = append(slice, "k")
	slice = append(slice, "a")
	slice = append(slice, "k")
	slice = append(slice, "a")
	slice = append(slice, "k")
	slice = append(slice, "a")
=======
=======
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
	var slice []int
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	slice = append(slice, 4)
	slice = append(slice, 5)
	slice = append(slice, 6)
	slice = append(slice, 7)
	slice = append(slice, 8)
	slice = append(slice, 9)
	slice = append(slice, 10)
<<<<<<< HEAD
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
=======
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
	fmt.Println(slice[0])
	fmt.Println(slice[9])

	// Tapi ada tips nih untuk membaca slice ataupun array. kita bisa menggunakan
	// for range yang akan menampilkan index dan nilai yang ada pada array tersebut.
	// for loop akan dijelaskan lebih lengkap di materi looping

	for index, value := range slice {
		fmt.Println(index, value)
	}
}
