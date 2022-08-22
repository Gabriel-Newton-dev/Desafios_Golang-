// Escreva um programa para ler 3 valores inteiros (considere que não
// 	serão lidos valores iguais) e escrevê-los em ordem crescente.

package main

import "fmt"

var A, B, C int

func main() {

	fmt.Print("Entre com o primeiro valor: ")
	fmt.Scan(&A)

	fmt.Print("Entre com o segundo valor: ")
	fmt.Scan(&B)

	fmt.Print("Entre com o terceiro valor: ")
	fmt.Scan(&C)

	if (A == B) && (A == C) && (B == C) {
		fmt.Println("Todos os números são iguais")
	} else if (A < B) && (B < C) && (C > A) && (C > B) {
		fmt.Println(A, B, C)
	} else if (A < B) && (A < C) && (B > A) && (B > C) {
		fmt.Println(A, C, B)
	} else if (A > B) && (A < C) && (B < C) && (C > A) {
		fmt.Println(B, A, C)
	} else if (C < A) && (C < B) && (B > C) && (B > A) {
		fmt.Println(C, A, B)
	} else if (C < A) && (C < B) && (B < A) && (B > C) {
		fmt.Println(C, B, A)
	} else if (A > B) && (A > C) && (B < C) && (C < A) {
		fmt.Println(B, C, A)
	}
}
