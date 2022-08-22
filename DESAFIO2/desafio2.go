// //2. Escreva  um  programa  para  ler  o  ano  de  nascimento  de  uma  pessoa  e  escrever
// uma  mensagem  que  diga  se  ela  poderá  ou  não  votar  este  ano  (não é necessário considerar o mês em que ela nasceu).

package main

import "fmt"

var digiteIdade int

func main() {

	maiorIdade := 16

	fmt.Print("Digite por favor ano que você nasceu: ")
	fmt.Scan(&digiteIdade)

	calcularIdade := 2022 - digiteIdade

	if calcularIdade >= maiorIdade {
		fmt.Printf("Sua idade atual é %v anos, sendo assim, você é apto para votar. # vote com consciência.", calcularIdade)
	} else {
		fmt.Printf("Sua idade atual é %v anos, sendo assim, você é inapto para votar, visto que, é menor que 16 anos. :( ", calcularIdade)
	}

}
