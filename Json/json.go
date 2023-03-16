package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

func main() {

	group1 := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Red", "Ruby", "Maroon"},
	}
	fmt.Println(group1)

	group1Json, err := json.Marshal(group1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(group1Json))

}
