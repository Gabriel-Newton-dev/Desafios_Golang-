package main

import (
	"encoding/json"
	"fmt"
)

type Students struct {
	Name string
	Age  int
}

func main() {

	Maria := Students{"Maria Izabel", 19}
	fmt.Println(Maria)

	MariaJson, err := json.Marshal(Maria)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(MariaJson))

}
