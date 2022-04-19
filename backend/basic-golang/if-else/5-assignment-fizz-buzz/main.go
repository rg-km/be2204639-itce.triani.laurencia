package main

import "fmt"

// Quiz ini dinamakan fizzbuzz
// Tuliskan program for loop dari 1 sampai dengan 100.
// lalu setiap perulangan program ini.

// Jika angka tersebut habis dibagi 3, maka tampilkan "fizz"
// Jika angka tersebut habis dibagi 5, maka tampilkan "buzz"
// Jika angka tersebut habis dibagi 3 dan 5, maka tampilkan "fizzbuzz"

// Tips gunakan % (modulo) untuk mengetahui apakah angka tersebut habis dibagi 3 atau 5 atau tidak.

func main() {
	for i := 1; i <= 100; i++ {
		// TODO: answer here
<<<<<<< HEAD
<<<<<<< HEAD
		if i % 3 == 0{
			fmt.Println(i,"fizz")
			
		}
		if i % 5 == 0{
			fmt.Println(i,"buzz")
		}
		if i % 3 == 0 && i % 5 == 0  {
			fmt.Println(i,"fizzbuzz")
		}
		fmt.Println(i)
	
=======
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
=======
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
	}
}
