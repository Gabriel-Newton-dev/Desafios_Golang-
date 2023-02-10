// - Crie um map com key tipo string e value tipo []string.
//     - Key deve conter nomes no formato sobrenome_nome
//     - Value deve conter os hobbies favoritos da pessoa
// - Demonstre todos esses valores e seus índices.

package main

import "fmt"

func main() {

	x := map[string][]string{
		"Gabriel": {"Programar", "viajar", "fazer musculação"},
		"Deise":   {"Viajar", "ver séries", "fazer compras"},
	}

	for key, value := range x {
		fmt.Println(key)
		for i, h := range value {
			fmt.Println("\t", i, "-", h)
		}
	}
}
