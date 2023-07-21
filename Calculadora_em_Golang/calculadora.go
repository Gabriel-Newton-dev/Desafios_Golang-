package main

import "fmt"

var number1, number2 float64

var operator string

func main() {

	fmt.Printf("Enter firts number: ")
	fmt.Scan(&number1)

	fmt.Printf("Enter second number: ")
	fmt.Scan(&number2)

	fmt.Printf("Enter operator ( + | - | * | / ): ")
	fmt.Scan(&operator)

	if operator == "+" {
		fmt.Printf("%v + %v = %v \n", number1, number2, number1+number2)
	}
	if operator == "-" {
		fmt.Printf("%v - %v = %v \n", number1, number2, number1-number2)
	}
	if operator == "*" {
		fmt.Printf("%v * %v = %v \n", number1, number2, number1*number2)
	}
	if operator == "/" {
		fmt.Printf("%v / %v = %v \n", number1, number2, number1/number2)
	}
}
