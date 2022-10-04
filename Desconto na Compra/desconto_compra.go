// Faça um programa que calcule 10% de deconto em cima da compra de um determinado equipamento.
// Para isso faça um metodo, para dar esses 10% de desconto.
// Faça uma struct e informe 2 valores diferentes.

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Notebooks struct {
	Marca       string
	Modelo      string
	Valor       int64
	Intel       bool
	Processador string
}

func (n Notebooks) desconto() {
	desc := (n.Valor * 10) / 100
	desc1 := n.Valor - desc
	fmt.Printf("O notebook %s está no valor de R$ %v reais, para pagamento a vista ele tem 10 porcento de desconto(%v), ficando um valor de R$ %v reais", n.Modelo, n.Valor, desc, desc1)
}

func main() {

	Dell := Notebooks{"Dell", "Xps", 5765, true, "I7"}
	//fmt.Printf("O Equipamento %s trata-se de um %s, no valor de R$ %v Reais.", Dell.Marca, Dell.Modelo, Dell.Valor)
	Dell.desconto()

	Macbook := Notebooks{"Apple", "Macbook", 9800, true, "I5"}
	Macbook.desconto()

	DellJson, err := json.Marshal(Dell)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(DellJson))

	MacJson, err := json.Marshal(Macbook)
	if err != nil {
		log.Fatal()
	}
	fmt.Println(string(MacJson))

}
