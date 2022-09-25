// Faça um programa que calcule a média do aluno, que seja possível colocar nome e sobrenome;
// Além das notas dos 4 bimestres, a media das notas se for acima de 7 aluno aprovado;
// Nota menor que 7 e maior que 4, aluno em recuperação, menor que 4 aluno reprovado.

package main

import "fmt"

var nome string
var sobrenome string
var bimestre1 float64
var bimestre2 float64
var bimestre3 float64
var bimestre4 float64

func main() {

	fmt.Print("Digite o nome do aluno: ")
	fmt.Scan(&nome)

	fmt.Print("Digite o sobrenome do aluno: ")
	fmt.Scan(&sobrenome)

	fmt.Print("Digite a nota do primeiro bimestre: ")
	fmt.Scan(&bimestre1)

	fmt.Print("Digite a nota do segundo bimestre: ")
	fmt.Scan(&bimestre2)

	fmt.Print("Digite a nota do terceiro bimestre: ")
	fmt.Scan(&bimestre3)

	fmt.Print("Digite a nota do quarto bimestre: ")
	fmt.Scan(&bimestre4)

	resultado(bimestre1, bimestre2, bimestre3, bimestre4)

}

func resultado(n1, n2, n3, n4 float64) {
	media := (n1 + n2 + n3 + n4) / 4
	if media >= 7.0 {
		fmt.Printf("Aluno %s %s está aprovado, com média de %v", nome, sobrenome, media)
	} else if media >= 4.0 && media < 7.0 {
		fmt.Printf("Aluno %s %s está em recuperação, com média de %v", nome, sobrenome, media)
	} else if media < 4.0 {
		fmt.Printf("Aluno %s %s está reprovado, ficom com a média %v.", nome, sobrenome, media)
	}
}
