// crie uma calculadora que possa receber numeros inteiros e possa usar os 4 operadores aritiméticos.

package main

import "fmt"

var number1 int
var number2 int
var operadores string

func main() {

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&number1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&number2)

	fmt.Print("Digite qual operador deseja utilizar ( +, -, * ou /) : ")
	fmt.Scan(&operadores)

	switch operadores {
	case "+":
		fmt.Println(number1 + number2)
	case "-":
		fmt.Println(number1 - number2)
	case "*":
		fmt.Println(number1 * number2)
	case "/":
		fmt.Println(number1 / number2)

	}
}
