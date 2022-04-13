package main

import "fmt"

func main() {
	// fungsi ini akan mengembalikan fungsi berikut
	// func(x, y int) int {
	// 	return x * y
	// }
	// TODO: answer here
<<<<<<< HEAD
	getAreaFunc := func (y int) func (int)int  {
		return func(x int) int {
			return x*y
		}
	}
	areaF := getAreaFunc(4)
	res := areaF(3) // 12
=======
	areaF := getAreaFunc()
	res := areaF(3, 4) // 12
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
	fmt.Println(res)
}
