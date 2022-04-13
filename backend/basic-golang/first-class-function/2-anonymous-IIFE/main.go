package main

import "fmt"

func main() {
	// fungsi yang lansung dijalankan
	func() {
		fmt.Println("Halo dari anonymous function")
	}()

	//contoh lain untuk melakukan pemangkatan 2
	func(x int) {
		fmt.Println(x * x)
<<<<<<< HEAD
	}(4)
=======
	}(3)
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c

}
