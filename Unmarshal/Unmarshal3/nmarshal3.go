package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Motorbike struct {
	Brand  string
	Model  string
	Value  float64
	Sporty bool
}

func main() {

	R := []byte(`{"Brand": "Yamaha","Model": "R1","Value": 125.900,"Sporty": true}`)

	var Yamaha Motorbike
	err := json.Unmarshal(R, &Yamaha)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(Yamaha)
	HondaUnmarshal()
	DafraMotorbike()
}

func HondaUnmarshal() {

	Hornet := []byte(`{"Brand": "Honda", "Model": "Hornet", "value": 49.850,"Sporty": true}`)

	var Honda Motorbike
	err := json.Unmarshal(Hornet, &Honda)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(Honda)

}

func DafraMotorbike() {
	
	Dafra := []byte(`{"Brand": "Dafra", "Model": "X5", "value": 15.302, "Sporty": false}`)

	var DafraUnmarchal Motorbike
	err := json.Unmarshal(Dafra, &DafraUnmarchal)
	if err != nil {
		fmt.Println(DafraUnmarchal)
	}
}
