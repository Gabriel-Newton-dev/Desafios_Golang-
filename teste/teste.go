// Create a function that takes the age and return the age in days.

// Examples
// calcAge(65) ➞ 23725

// calcAge(0) ➞ 0

// calcAge(20) ➞ 7300

//expect only positive integer inputs.
package main

import "fmt"

var idade int
var resultado int

func main() {
	fmt.Println("Digite a sua idade por favor:")
	fmt.Scan(&idade)
	ConversorDeIdadeEmDias()
}

func ConversorDeIdadeEmDias() {
	if idade <= 0 {
		fmt.Println("Idade Inválida.")
	} else {
		fmt.Println("A sua idade convertida em dias é de:", idade*365, "dias.")
	}
}
