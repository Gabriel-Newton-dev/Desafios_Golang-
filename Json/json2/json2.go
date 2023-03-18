package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name        string
	Surname     string
	Age         int
	Profession  string
	BankAccount float64
}

func main() {

	jamesBond := Person{"James", "Bond", 51, "Secret Agent", 5000000.560}
	fmt.Println(jamesBond)

	jamesBondJson, err := json.Marshal(&jamesBond)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jamesBondJson))

}
