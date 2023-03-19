package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name         string  `json:"name"`
	Surname      string  `json:"surname"`
	Age          int     `json:"age"`
	Profession   string  `json:"profession`
	BankAccounts float64 `json:"bankAccounts`
}

func main() {

	Json := []byte(`{"Name":"James","Surname":"Bond","Age":51,"Profession":"Secret Agent","BankAccount":5000000.56}`)

	RonaldinhoGaucho := []byte{"Name": "Ronaldinho", "Surname": "Gaúcho", "Age": 49, "Profession": "Jgador de futebol", "BankAccounts": "5000000"}

	var jameBond Person
	err := json.Unmarshal(Json, &jameBond)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(jameBond)

}
