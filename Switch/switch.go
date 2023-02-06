//- Crie um programa que utilize a declaração switch, onde o switch statement seja
//uma variável do tipo string com identificador "esporteFavorito".

package main

import "fmt"

var esporteFavorito string

func main() {

	fmt.Print("Adivinha qual o meu esporte favorito ( voley, natação ou futebol) ? ")
	fmt.Scan(&esporteFavorito)

	switch esporteFavorito {
	case "futebol":
		fmt.Println("Parabéns você acertou.")
	case "voley":
		fmt.Println("Parabéns você acertou.")
	case "natacao":
		fmt.Println("Parabéns você acertou.")
	default:
		fmt.Println("Não localizado no rol de esporte")
	}
}
