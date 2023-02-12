// - Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
//     - Nome
//     - Sobrenome
//     - Sabores favoritos de sorvete
// - Crie dois valores do tipo "pessoa" e demonstre estes valores,
// utilizando range na slice que contem os sabores de sorvete.

package main

import "fmt"

type pessoa struct {
	nome             string
	sobrenome        string
	sabores_sorvetes []string
}

func main() {

	gabriel := pessoa{
		nome:             "Gabriel",
		sobrenome:        "Newton",
		sabores_sorvetes: []string{"Napolitano", "Passas ao rum"},
	}

	fmt.Println(gabriel.nome, gabriel.sobrenome)
	for _, value := range gabriel.sabores_sorvetes {
		fmt.Println("\t", value)
	}

	deise := pessoa{
		nome:             "Deise",
		sobrenome:        "Santos",
		sabores_sorvetes: []string{"Flocos", "Passas ao rum"},
	}

	fmt.Println(deise.nome, deise.sobrenome)
	for _, value := range deise.sabores_sorvetes {
		fmt.Println("\t", value)
	}

}
