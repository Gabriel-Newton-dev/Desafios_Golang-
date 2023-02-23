// - Crie um map com key tipo string e value tipo []string.
//     - Key deve conter nomes no formato modelo do carro
//     - Value deve conter as especificacoes do carro.
// - Demonstre todos esses valores e seus índices.

package main

import "fmt"

type carros struct {
	marca          string
	modelo         string
	especificacoes []string
	turbo          bool
}

func main() {

	veiculo := make(map[string]carros)

	veiculo["Cruze"] = carros{"Chevrolet", "Cruze Ltz", []string{"mais completo da categoria,", "bancos de couro"}, true}

	veiculo["Audi"] = carros{"Audi", "A3", []string{"teto solar, ", "bancos em couro"}, true}

	veiculo["Eclipse"] = carros{"Mitsubish", "Eclipse", []string{"Teto solar,", "bancos elétricos em couro"}, true}

	for key, value := range veiculo {
		fmt.Println(key, value)
	}

}
