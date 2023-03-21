package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Cars struct {
	Name  string
	Model string
	Value float64
	Turbo bool
}

func main() {

	Eclipse := []byte(`{"Eclipse", "GST", 49.850, true}`)

	var mitsubish Cars
	err := json.Unmarshal(Eclipse, &mitsubish)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mitsubish)

	ChevroletUnmarshal()

}

func ChevroletUnmarshal() {
	Montana := []byte(`{"Montana", "LTZ", 98.760, true}`)

	var Chevrolet Cars
	err := json.Unmarshal(Montana, &Chevrolet)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(Chevrolet)
}
