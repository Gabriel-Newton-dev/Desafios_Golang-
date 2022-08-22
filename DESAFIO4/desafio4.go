// As maçãs  custam  R$  0,30  cada  se  forem  compradas menos  do  que  uma  dúzia,  e  R$  0,25  se  forem  compradas
// pelo  menos  doze.  Escreva  um  programa  que  leia  o  número  de  maçãs  compradas,  calcule  e  escreva  o  valor
// total da compra.

package main

import "fmt"

var vendaDeMaca int

func main() {

	promocao := 12

	semDesconto := 0.30
	comDesconto := 0.25

	fmt.Print("Pois não, quantas maças deseja comprar: ")
	fmt.Scan(&vendaDeMaca)

	VendaMaiorDoze := float64(vendaDeMaca) * comDesconto
	vendaMenorDoze := float64(vendaDeMaca) * semDesconto

	if vendaDeMaca >= promocao {
		fmt.Printf("Você comprou %v maças, valor com desconto R$ 0.25 centavos, valor total da compra: %v Reais", vendaDeMaca, VendaMaiorDoze)
	} else {
		fmt.Printf("Você comprou %v maças, saindo R$ 0.30 centavos cada, valor total da compra : %v", vendaDeMaca, vendaMenorDoze)
	}

}
