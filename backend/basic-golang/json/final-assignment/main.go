package main

import (
	"encoding/json"
<<<<<<< HEAD
<<<<<<< HEAD
)

// TODO: answer here
type Ruang struct {
	RuangTamu Items `json:"ruangTamu"`
}

type Item struct {
	Nama   string `json:"nama"`
	Jumlah int    `json:"jumlah"`
	Warna  string `json:"warna"`
	Ukuran Ukuran `json:"ukuran"`
}

type Items struct {
	Items []Item `json:"items"`
}

type Ukuran struct {
	Panjang string `json:"panjang"`
	Tinggi  string `json:"tinggi"`
}

func (r Ruang) EncodeJSON() string {
	// TODO: answer here
	// encode
	data, _ := json.Marshal(r)
	// convert to string
	return string(data)

=======
=======
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
	"log"
)

// buat JSON string array nested seperti berikut
/*
{
	"ruangTamu": {
			"items": [
					{
							"nama": "Meja",
							"jumlah": 20,
							"warna": "Coklat",
							"ukuran": {
									"panjang": "50 cm",
									"tinggi": "25 cm"
							}
					},
					{
							"nama": "Kursi",
							"jumlah": 1,
							"warna": "Hitam",
							"ukuran": {
									"panjang": "70 cm",
									"tinggi": "30 cm"
							}
					}
			]
	}
}
*/

// TODO: answer here

func (r Ruang) EncodeJSON() string {
	// TODO: answer here
<<<<<<< HEAD
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
=======
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
}

func NewRuang(r Ruang) Ruang {
	return r
}
<<<<<<< HEAD
<<<<<<< HEAD

func main() {

	// masukan data ke struct Item
	mejaMeja := []Item{
		
		{
			Nama:   "Meja",
			Jumlah: 20,
			Warna:  "Coklat",
			Ukuran: Ukuran{
				Panjang: "50 cm",
				Tinggi:  "25 cm",
			},
		},
		{
			Nama:   "Kursi",
			Jumlah: 1,
			Warna:  "Hitam",
			Ukuran: Ukuran{
				Panjang: "70 cm",
				Tinggi:  "30 cm",
			},
		},
	
	}

	// masukan mejaMeja ke struct Items
	items := Items{
		mejaMeja,
	}

	ruang := Ruang{
		RuangTamu: items,
	}

	// encode ke funtion EncodeJSON
	encode := ruang.EncodeJSON()

	// print
	println(encode)
}
=======
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
=======
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
