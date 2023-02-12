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
	fmt.Println(mape)

}
