// - Crie uma slice contendo slices de strings ([][]string). Atribua valores a
// este slice multi-dimensional da seguinte maneira:
//     - "Nome", "Sobrenome", "Hobby favorito"
// - Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.

package main

import "fmt"

func main() {

	sliceString := [][]string{
		{"Gabriel", "Santos", "Programar"},
		{"Larissa", "Santos", "Viajar"},
		{"Ana", "Madalena", "Voley"},
	}

	for _, value := range sliceString {
		fmt.Println(value)
	}

}

// type Persons struct {
// 	Nome          string
// 	Sobrenome     string
// 	HobbyFavorito string
// }

// var sliceString []Persons

// func main() {

// 	sliceString = []Persons{{"Gabriel", "Santos", "Programar"},
// 		{"Larissa", "Santos", "Viajar"}, {"Ana", "Madalena", "Voley"}}

// 	for i := 0; i < len(sliceString); i++ {
// 		fmt.Println(sliceString[i])
// 	}

// no caso se for Range

// for _, value := range sliceString{
// 	fmt.Println(value)
// }

// }
