package main

import "fmt"

func main() {
	//membuat rectangle dengan anonymous struct
	//field dari struct ini sama seperti rectangle sebelumnya
	// TODO: answer here
<<<<<<< HEAD
	newRectangle := struct {
		Width int
		Length  int
	}{
		Width: 10,
		Length:  20,
	}
=======
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c

	fmt.Printf("rectangle dengan lebar %d dan panjang %d", newRectangle.Width, newRectangle.Length)
}
