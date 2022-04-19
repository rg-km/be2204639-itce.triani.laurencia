package main

import (
	"errors"
	"fmt"
)

// Dari contoh yang telah diberikan, kalian dapat mencoba membuat custom error baru dengan atribut message dan errCode.
// Misalnya adalah error untuk validasi data umur kurang dari 0.

// TODO: answer here
<<<<<<< HEAD
type ErrorUmurTidakAda struct {
	message string
	errCode int32
}

func (e *ErrorUmurTidakAda) Error() string {
	return fmt.Sprintf("error invalid data %d: %s", e.errCode, e.message)
}
=======
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015

func GetAge(data map[string]int, name string) (int, error) {
	if _, ok := data[name]; !ok {
		return 0, errors.New("Data not found")
	}

	if data[name] < 0 {
		// Isilah baris ini dengan return 0 dan custom error yang telah dibuat dengan message error invalid data dan errCode 500
		// TODO: answer here
<<<<<<< HEAD
		return 0, &ErrorUmurTidakAda{
			message: fmt.Sprintf("%s not found in data", name),
			errCode: 500,
		}
=======
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
	}

	return data[name], nil
}

func main() {
	peopleAge := map[string]int{
		"John": 20,
		"Raz":  8,
		"Tony": -1,
	}

	_, err := GetAge(peopleAge, "Tony")
	if err != nil {
		fmt.Println(err.Error())
	}
}
