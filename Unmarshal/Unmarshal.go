package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Age        int    `json:"age"`
	Profession string `json:"profession"`
	//BankAccounts float64 `json:"bankAccounts"`
}

func main() {

	RonaldinhoGaucho := []byte(`{"Name": "Ronaldinho", "Surname": "Gaúcho", "Age": 49, "Profession": "Jogador de futebol"}`)

	var Ronaldinho Person
	err := json.Unmarshal(RonaldinhoGaucho, &Ronaldinho)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(Ronaldinho)

}
