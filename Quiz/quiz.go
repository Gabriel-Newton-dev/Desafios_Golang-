package main

import "fmt"

func lenTeste() {
	data := make([]int, 5)
	for i := 0; i < 3; i++ {
		data[i] = i
	}
	fmt.Println(len(data), cap(data), data)
}

func main() {
	lenTeste()
}
