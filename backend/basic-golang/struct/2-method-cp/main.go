package main

import "fmt"

// buat struct Rectangle (persegi panjang) dengan dua atribut yaitu Width dan Length
// tambah dua method :
// GetArea() dan GetPerimeter()
// GetArea() digunakan untuk menampilkan (print) luas persegi panjang
// GetPerimeter() digunakan untuk menampilkan (print) keliling persegi panjang

type Rectangle struct {
	// TODO: answer here
<<<<<<< HEAD
	Width int
	Length int
}

func (r Rectangle) GetArea(){
	Area := r.Width*r.Length
	fmt.Println("Area :", Area)
}

func (r Rectangle) GetPerimeter(){
	Parimeter := 2*(r.Width+r.Length)
	fmt.Println("Parimeter :", Parimeter)
=======
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
}

// TODO: answer here
func main() {
	var r Rectangle
	r.Width = 10
	r.Length = 20
	fmt.Println(r)

	r.GetArea()
	r.GetPerimeter()
}
