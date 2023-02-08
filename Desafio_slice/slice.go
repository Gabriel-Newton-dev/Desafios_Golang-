// - Crie uma slice de tipo int
// - Atribua 10 valores a ela
// - Utilize range para demonstrar todos estes valores.
// - E utilize format printing para demonstrar seu tipo.

// PARTE 2 DO DESAFIO

// Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:

package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, value := range slice {
		fmt.Println(value)
	}

	fmt.Printf("%T", slice)
	// - Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
	fmt.Println(slice[:3])
	// - Do quinto ao último item do slice (incluindo o último item!)
	fmt.Println(slice[4:])
	// - Do segundo ao sétimo item do slice (incluindo o sétimo item!)
	fmt.Println(slice[1:7])
	// - Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
	fmt.Println(slice[2:9])
	// - Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item
	fmt.Println(slice[2 : len(slice)-1])

}
