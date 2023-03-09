// - Crie um tipo struct "pessoa" que contenha os campos:
//     - nome
//     - sobrenome
//     - idade
// - Crie um método para "pessoa" que demonstre o nome completo e a idade;
// - Crie um valor de tipo "pessoa";
// - Utilize o método criado para demonstrar esse valor.

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name    string
	Surname string
	Age     int
}

// metodo criado com nome Persons
func (p *Person) Persons() {
	fmt.Println(p.Name, p.Surname, p.Age)

}

func main() {

	Maria := Person{"Maria", "Izabel", 29}
	Maria.Persons()

	Gustavo := Person{"Gustavo", "Lopes", 32}
	Gustavo.Persons()

	MariaJson, err := json.Marshal(Maria)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(MariaJson))

	GustavoJson, err := json.Marshal(Gustavo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(GustavoJson))
}
