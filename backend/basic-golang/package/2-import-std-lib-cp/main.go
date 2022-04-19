package main

import (
	// TODO: answer here
<<<<<<< HEAD
<<<<<<< HEAD
	"fmt"
	"time"
=======
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
=======
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
)

// Dari contoh yang telah diberikan dan eksplorasi yang dilakukan dari standard library golang, kamu dapat mencoba untuk mengimport salah satu package pada golang.
// Cobalah untuk mengimport package time dan lakukan perhitungan perbedaan hari antara dua waktu.

func CountDays(start, end time.Time) int {
	// TODO: answer here
<<<<<<< HEAD
<<<<<<< HEAD
	duration := end.Sub(start).Hours()
	days := duration / 24
	return int(days)
=======
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
=======
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
}

func main() {
	start := time.Date(2022, time.March, 21, 0, 0, 0, 0, time.UTC)
	end := time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC)
	dayDifference := CountDays(start, end)
	fmt.Println(dayDifference)
<<<<<<< HEAD
<<<<<<< HEAD
}
=======
}
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
=======
}
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
