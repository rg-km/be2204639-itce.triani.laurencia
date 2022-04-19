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

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan condition pada template.
// Lengkapi function ExecuteToByteBuffer dan variabel textTemplate sehingga memiliki keluaran:
// 1. Jika saldo sama dengan 0, cetak "Akun Tony dengan Nomor 1002321 tidak memiliki saldo."
// 2. Jika saldo lebih dari 0, cetak "Akun Tony dengan Nomor 1002321 memiliki saldo sebesar $1000."

type Account struct {
	Name    string
	Number  string
	Balance int
}

func ExecuteToByteBuffer(account Account) ([]byte, error) {
<<<<<<< HEAD
<<<<<<< HEAD

	// buat kondisi
	// jika saldo sama dengan 0, cetak "Akun Tony dengan Nomor 1002321 tidak memiliki saldo."
	// jika saldo lebih dari 0, cetak "Akun Tony dengan Nomor 1002321 memiliki saldo sebesar $1000."
	// gt adalah operator untuk mengecek nilai yang lebih besar dari nilai yang diberikan.
	// singkatan gt adalah greater than.
	// jika kita ingin mengecek nilai yang lebih kecil dari nilai yang diberikan, maka gunakan singkatan lt.
	// singkatan lt adalah less than.
	
	textTemplate := `{{if (gt .Balance 0) }}Akun {{ .Name }} dengan Nomor {{ .Number }} memiliki saldo sebesar ${{ .Balance }}.{{ else }}Akun {{ .Name }} dengan Nomor {{ .Number }} tidak memiliki saldo.{{ end }}`

	// bikin template
	tmpl, err := template.New("test").Parse(textTemplate)
	if err != nil {
		return nil, err
	}

	// bikin buffer
	var b bytes.Buffer

	// eksekusi template
	err = tmpl.Execute(&b, account)
	if err != nil {
		return nil, err
	}

	// kembalikan hasil
	return b.Bytes(), nil

}

func main() {
	// isi data Account
	account := Account{
		Name:    "Tony",
		Number:  "1002321",
		Balance: 0,
	}

	// eksekusi ExecuteToByteBuffer
	b, err := ExecuteToByteBuffer(account)
	if err != nil {
		panic(err)
	}

	// cetak hasil
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
