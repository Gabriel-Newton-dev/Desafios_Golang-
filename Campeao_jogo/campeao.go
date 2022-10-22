// Gols das partidas - time A [2,5,0,4]
// Gols do time B [1,0,2,1]
// em cada partida informe quem foi o time ganhador de cada partida

package main

import "fmt"

func main() {

	timeA := [3]int{2, 5, 2}
	timeB := [3]int{1, 0, 2}

	resultadoTimeA := 0

	switch {
	case timeA[1] > timeB[0]:
		fmt.Println(resultadoTimeA + 1)
	case timeA[1] > timeB[01]:
		fmt.Println(resultadoTimeA + 1)
	case timeA[2] > timeB[2]:
		fmt.Println(resultadoTimeA + 1)
	case timeA[0] < timeB[0]:
		fmt.Println("Time B fez 1 ponto")
	case timeA[1] < timeB[1]:
		fmt.Println("Time B fez 1 ponto")
	case timeA[2] < timeB[2]:
		fmt.Println("Time B fez 1 ponto")

	case timeA[0] == timeB[0] && timeA[1] == timeB[1] && timeA[2] == timeB[2]:
		fmt.Println("Os times empataram")

	}

}
