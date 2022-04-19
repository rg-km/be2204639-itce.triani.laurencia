package main

import "fmt"

func main() {
	//fungsi counter akan menerima (x int) dan mengembalikan fungsi
	//fungsi yang dikembalikan akan melakukan decrement nilai parameter x ketika dipanggil dan
	//mengembalikan nilai parameter x

	counter := func(x int) func() int {
		// TODO: answer here
<<<<<<< HEAD
<<<<<<< HEAD
		return func() int {
			x++
			return x
		}
=======
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
=======
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
	}
	decrement := counter(5)
	fmt.Println(decrement())
	fmt.Println(decrement())
	fmt.Println(decrement())
}
