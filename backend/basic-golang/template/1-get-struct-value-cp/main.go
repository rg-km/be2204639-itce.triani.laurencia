package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
)

//Buat struct Student dengan field Name tipe data <string> dan ScoreAverage tipe data <float64>
//tampilkan dengan wording:
//Hello Rogu,
//Nilai rata rata kamu 7.8

type Student struct {
	// TODO: answer here
<<<<<<< HEAD
	Name         string
	ScoreAverage float64
=======
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
}

// main function
func main() {
	buff := new(bytes.Buffer)
	// TODO: answer here
<<<<<<< HEAD
	std := Student{
		Name:         "Rogu",
		ScoreAverage: 7.8,
	}

	// bikin tempalte
	tmp1, _ := template.New("student").Parse("Hello {{.Name}}, \nNilai rata rata kamu {{.ScoreAverage}}")
=======
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c

	if err := tmp1.Execute(buff, std); err != nil {
		log.Fatalf("execute template error: %s", err.Error())
	}
	fmt.Println(buff.String())
}
