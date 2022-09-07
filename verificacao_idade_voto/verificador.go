// Escreva  um  programa  para  ler  o  ano  de  nascimento  de  uma  pessoa  e  escrever
// uma  mensagem  que  diga  se  ela  poderá  ou  não  votar  este  ano  (não é necessário
//considerar o mês em que ela nasceu).

package main

import "fmt"

var anoDenascimento int64

func main() {

	fmt.Printf("Digite o ano do seu nascimento: ")
	fmt.Scan(&anoDenascimento)

	verificador := 2022 - anoDenascimento

	if verificador >= 16 {
		fmt.Printf("Você tem %v idades, sendo assim está apto para votar nesse ano", verificador)
	}
	if verificador < 16 {
		fmt.Printf("Você tem %v idades, sendo assim está inapto para partiticipar da votação", verificador)
	}

}
