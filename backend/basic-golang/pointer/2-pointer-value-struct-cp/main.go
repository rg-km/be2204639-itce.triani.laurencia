package main

<<<<<<< HEAD
<<<<<<< HEAD
import (
	"fmt"

	// "golang.org/x/text/width"
)
=======
import "fmt"
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
=======
import "fmt"
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015

//buat struct Rectangle dengan dua atribut yaitu Width dan Length
// tambah dua method :
// SetWidthPointer(width int) dan SetWidthValue(width int)
// SetWidthPointer(width int) untuk mengubah width dengan pointer receiver
// SetWidthValue(width int) untuk mengubah width dengan value

// TODO: answer here
<<<<<<< HEAD
<<<<<<< HEAD
type Rectangle struct {
	width int
	Length int
}
func (r *Rectangle) SetWidthPointer(width int){
	fmt.Println("hasil :", r.width + width) 
}
func (r Rectangle) SetWidthValue(width int) {
	fmt.Println("hasil :", r.width + width) 
}
func main() {
	var r Rectangle
	r.width = 10
	r.Length = 20

	fmt.Println("sebelum melakukan set width dengan pointer", r.width)
	r.SetWidthPointer(30)
	fmt.Println("sesudah melakukan set width dengan pointer", r.width)
	r.SetWidthValue(70)
=======
=======
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
func main() {
	var r Rectangle
	r.Width = 10
	r.Length = 20

	fmt.Println("sebelum melakukan set width dengan pointer", r.Width)
	r.SetWidthPointer(30)
	fmt.Println("sesudah melakukan set width dengan pointer", r.Width)
	r.SetWidthValue(70)
	fmt.Println("sesudah melakukan set width dengan value", r.Width)
<<<<<<< HEAD
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
=======
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
}
