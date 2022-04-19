package main

import (
	"bytes"
<<<<<<< HEAD
<<<<<<< HEAD
	"fmt"
=======
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
=======
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
	"html/template"
)

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan looping pada template.
// Lengkapi function ExecuteToByteBuffer dan textTemplate sehingga template dapat mencetak objek Leaderboard dengan ketentuan:
// Lakukan looping untuk setiap user
// Pada setiap loop, cetak "Peringkat ke-n: Nama", contoh: "Peringkat ke-1: Roger"

type UserRank struct {
	Name  string
	Email string
	Rank  int
}

type Leaderboard struct {
	Users []*UserRank
}

func ExecuteToByteBuffer(leaderboard Leaderboard) ([]byte, error) {
<<<<<<< HEAD
<<<<<<< HEAD

	// bikin kalimat di dalam looping
	// range looping berfungsi untuk mengambil data dari slice
	textTemplate := `{{range .Users}}Peringkat ke-{{.Rank}}: {{.Name}} {{end}}`

	// bikin template
	tmpl, err := template.New("test").Parse(textTemplate)
	if err != nil {
		return nil, err
	}

	// bikin buffer
	var b bytes.Buffer

	// generate template
	err = tmpl.Execute(&b, leaderboard)
	if err != nil {
		return nil, err
	}

	// kembalikan hasil
	return b.Bytes(), nil

}

func main() {
	// isi data ke dalam struct
	users := []*UserRank{
		{
			Name: "Roger",
			Rank: 1,
		},
		{
			Name: "Tony",
			Rank: 2,
		},
		{
			Name: "Bruce",
			Rank: 3,
		},
		{
			Name: "Natasha",
			Rank: 4,
		},
		{
			Name: "Clint",
			Rank: 5,
		},
	}

	// masukan data Leaderboard
	leaderboardObject := Leaderboard{
		Users: users,
	}

	// generate output
	b, err := ExecuteToByteBuffer(leaderboardObject)
	if err != nil {
		panic(err)
	}

	// cetak output
	fmt.Println(string(b))

=======
	var textTemplate string
	// TODO: answer here
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
=======
	var textTemplate string
	// TODO: answer here
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
}
