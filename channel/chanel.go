package main

import (
	"fmt"
	"time"
)

func number1(done chan int) {
	for i := 0; i <= 10; i++ {
		done <- i
		fmt.Printf("Escrito no chanel: %d - ", i)
		time.Sleep(time.Millisecond * 150)
	}
	fmt.Println("Fim da escrita !")
	close(done)

}

func main() {

	// 3 no channel é o buffer.
	cn1 := make(chan int, 3)
	go number1(cn1)

	for v := range cn1 {
		fmt.Printf("Lido do channel: %d \n", v)
		time.Sleep(time.Millisecond * 130)
	}

}
