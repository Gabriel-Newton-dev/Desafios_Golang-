// Caso 1: Jogo de futebol com 2 times A e B, decidir qual time é vencedor
// golA = 2
// golB = 1

package main

import "fmt"

func main() {

	golA := 2
	golB := 1

	if golA > golB {
		fmt.Println("GolA é o vencedor do jogo")
	} else if golA < golB {
		fmt.Println("GolB é o vencedor do jogo")
	} else {
		fmt.Println("Houve empate")
	}
}
