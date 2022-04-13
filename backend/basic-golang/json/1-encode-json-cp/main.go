package main

import (
	"encoding/json"
<<<<<<< HEAD
	"fmt"
=======
	"log"
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
)

// Dari struct dan nama field yang sama dari contoh
// Buat string JSON dengan hasil seperti berikut
// {"jenis":"Meja Belajar","color":"green","jumlah":2}

type Meja struct {
	// TODO: answer here
<<<<<<< HEAD
	Jenis  string `json:"jenis"`
	Warna  string `json:"color"`
	Jumlah int    `json:"jumlah"`
=======
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
}

func (m Meja) EncodeJSON() string {
	// TODO: answer here
<<<<<<< HEAD
	// konversti sturct ke json
	data, _ := json.Marshal(m)
	return string(data)

=======
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
}

func NewMeja(m Meja) Meja {
	return m
}
<<<<<<< HEAD

func main() {
	// masukan data ke struct
	var meja Meja
	meja.Jenis = "Meja Belajar"
	meja.Warna = "green"
	meja.Jumlah = 2

	// isi data struct ke leaderboard
	result := meja.EncodeJSON()

	// cek hasil
	fmt.Println(result)
	// fmt.Println(meja.EncodeJSON())

}
=======
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
