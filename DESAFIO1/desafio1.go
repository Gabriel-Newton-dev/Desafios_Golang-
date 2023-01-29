// // Create a function that takes the age and return the age in days.
// Notes
// Use 365 days as the length of a year for this challenge.
// Ignore leap years and days between last birthday and now.
// Expect only positive integer inputs

package main

import "fmt"

var anoAtual int
var anoDoNascimento int

func main() {

	fmt.Println("Seja bem-vindo ao programa que converte a sua idade em dias!!")
	fmt.Print("Por favor entre com ano do seu nascimento: ")
	fmt.Scan(&anoDoNascimento)
	ReturnAgeDays()

}

func ReturnAgeDays() {
	anoAtual = 2023
	idade := anoAtual - anoDoNascimento
	resultado := idade * 365
	if anoAtual == anoDoNascimento {
		fmt.Println("Você ainda não completou 1 ano de idade")
	} else {
		fmt.Printf("A sua idade atuel é %d idade, e convertida em dias %d dias", idade, resultado)
	}
}

// o gabarito desse desafio está abaixo, porém achei uma forma muito mais simplificada de fazer o programa, trazendo o mesmo resultado.

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// const (
// 	daysInYear = 365
// )

// func main() {
// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Print("Enter your age: ")

// 	input, _ := reader.ReadString('\n')

// 	ageResult, err := CalculateAge(input)

// 	if err != nil {
// 		fmt.Printf("Sorry, but something went wrong: %v", err)
// 	} else {
// 		fmt.Printf("You are %d days old \n", ageResult)
// 	}

// }

// func CalculateAge(input string) (int64, error) {

// 	trimmedInput := strings.TrimRight(input, "\r\n")
// 	parsedAge, err := strconv.ParseInt(trimmedInput, 10, 64)
// 	if err != nil {
// 		return parsedAge, err
// 	}

// 	ageinDays := parsedAge * daysInYear

// 	return ageinDays, err
// }
