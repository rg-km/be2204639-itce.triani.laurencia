// Reverse

package main

import "fmt"

func Reverse(st []string, depth int) string {
	str := ""
	// TODO: answer here
	// loop untuk mengambil nilai dari array
	for i := depth; i >= 0; i-- {
		str += st[i]
	}
	return str
}

func main() {
	str := []string{"A", "I", "S", "E", "N", "O", "D", "N", "I"}
	s := Reverse(str, len(str)-1)
	fmt.Println(s)
}
