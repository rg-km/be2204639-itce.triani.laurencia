package main

<<<<<<< HEAD
import (
	"fmt"
)
=======
import "fmt"
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c

//buat struct Rectangle dengan dua atribut yaitu Width dan Length

// TODO: answer here
<<<<<<< HEAD
type Rectangle struct{
	Width int
	Length int
}
=======
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c

func main() {
	var r Rectangle
	r.Width = 10
	r.Length = 20
	fmt.Println(r) //{10,20}

	r2 := Rectangle{Width: 5, Length: 15}
	fmt.Println(r2) // {5,15}

	fmt.Println("lebar rectangle2:", r2.Width)    //5
	fmt.Println("panjang rectangle2:", r2.Length) //15

}
