package main

import "fmt"

func receiveBlock(output chan int) {
	c := make(chan int)
	result := 0
	go func() {
		fmt.Println("send to channel c")
		//kirim 1 ke channel c
		// TODO: answer here
		c <- 1
		fmt.Println("send to channel output")
	}()

	//result menerima data dari channel c
	// TODO: answer here
	result = <-c
	fmt.Println("receive from channel c")

	output <- result
	fmt.Println(c) //agar variabel c digunakan
}
