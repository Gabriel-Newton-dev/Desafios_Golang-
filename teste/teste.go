package main

import "fmt"

var number1, number2 float64
var operator string

func main() {

	fmt.Print("Enter the first number: ")
	fmt.Scan(&number1)

	fmt.Print("Enter the Second number: ")
	fmt.Scan(&number2)

	fmt.Print("Enter the operator ( + | - | * | /): ")
	fmt.Scan(&operator)

	calculator()

}

func calculator() {
	if number1 >= 0 && number2 >= 0 {
		switch operator {
		case "+":
			fmt.Printf("%v + %v = %v\n", number1, number2, number1+number2)
		case "-":
			fmt.Printf("%v - %v = %v\n", number1, number2, number1-number2)
		case "*":
			fmt.Printf("%v * %v = %v\n", number1, number2, number1*number2)
		case "/":
			fmt.Printf("%v / %v = %v\n", number1, number2, number1/number2)
		default:
			fmt.Println("please enter a valid operator")
		}
	} else {
		fmt.Println("please enter a valid number")
	}
}
