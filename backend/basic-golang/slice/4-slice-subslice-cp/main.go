package main

import "fmt"

// Disini kita akan mencoba untuk melakukan subslicing pada slice.
// Coba langsung gunakan function append ketika melakukan subslicing.
// contoh slice = append (slice, slice2[0:3])

// Silahkan copy slice dan mempunyai value "Marcus", "is", "known", "to", "be" dan "a", "philosopher"
// Silahkan print slice tersebut
// Expected output : [Marcus is known to be a philosopher]
func main() {
	slice := []string{"Marcus", "is", "known", "to", "be", "one", "of", "five", "greatest", "emperors", "of", "rome",
		"Aurelius", "is", "also", "known", "to", "be", "a", "philosopher"}

	// TODO: answer here
<<<<<<< HEAD
<<<<<<< HEAD
	var slice2 []string

	slice2 = append(slice2,slice... )
	s0 := slice2[:7]
	s0[5] = "a"
	s0[6] = "philosopher"
	fmt.Println(slice2)
	fmt.Println(s0)
=======
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
=======
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
}
