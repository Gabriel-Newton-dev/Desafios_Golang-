// - Crie um map com key tipo string e value tipo []string.
//     - Key deve conter nomes no formato sobrenome_nome
//     - Value deve conter os hobbies favoritos da pessoa
// - Demonstre todos esses valores e seus índices.

package main

import "fmt"

func main() {

	mape := map[string][]string{
		"Gabriel Newton": {"Desenvolver", "Estudar"},
		"Deise Santos":   {"Ver Filmes", "Viajar"},
		"Caio Mello":     {"Jogar PS5", "Brincar"},
		"Hulk":           {"Passear", "Comer ração"},
	}

	for key, value := range mape {
		fmt.Println(key)
		for chave, valor := range value {
			fmt.Println("\t", chave, " - ", valor)
		}
	}

}
