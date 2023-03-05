// - Crie uma função que receba um parâmetro variádico do tipo int e
//retorne a soma de todos os ints recebidos;
// - Passe um valor do tipo slice of int como argumento para a função;
// - Crie outra função, esta deve receber um valor do tipo slice of int e
//retornar a soma de todos os elementos da slice;
// - Passe um valor do tipo slice of int como argumento para a função.

package main

import "fmt"

func main() {

	sliceInt := []int{90, 100, 200, 500}
	sliceInt2 := []int{10, 20, 30, 40, 50}

	fmt.Println(variadico(sliceInt...))
	fmt.Println(somaSlice(sliceInt2))

}

func variadico(x ...int) int {
	total := 0
	for _, value := range x {
		total += value
	}
	return total
}

func somaSlice(y []int) int {
	total := 0
	for _, value := range y {
		total += value
	}
	return total
}
