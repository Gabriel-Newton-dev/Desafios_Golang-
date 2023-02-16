// - Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map,
//utilizando os sobrenomes como key.
// - Demonstre os valores do map utilizando range.
// - Os diferentes sabores devem ser demonstrados utilizando outro range,
//dentro do range anterior.

package main

import "fmt"

type Pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {
	//.............   KEY   VALUE
	mape := make(map[string]Pessoa)

	mape["VALADARES"] = Pessoa{
		nome:      "Jailton",
		sobrenome: "Valadares",
		sabores:   []string{"Morango", "Chocolate"}}

	mape["JÚNIOR"] = Pessoa{"Wladimir", "Júnior", []string{"Baunilha", "Passas ao Rum"}}

	for key, value := range mape {
		fmt.Println(key, value)
		for _, value := range value.sabores {
			fmt.Printf("Sorvete de %s \n", value)
		}

	}
}
