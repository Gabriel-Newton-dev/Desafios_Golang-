// // Create a function that takes the age and return the age in days.
// Notes
// Use 365 days as the length of a year for this challenge.
// Ignore leap years and days between last birthday and now.
// Expect only positive integer inputs

package main

import "fmt"

var idade int

func main() {

	fmt.Print("Digite a sua idade: ")
	fmt.Scan(&idade)

	conversor := 365 * idade

	fmt.Printf("A sua idade é %v anos, convertida em dias é: %v dias.", idade, conversor)

}

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
