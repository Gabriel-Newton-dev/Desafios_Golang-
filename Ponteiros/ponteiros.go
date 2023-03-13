// - Crie um valor e atribua-o a uma variável.
// - Demonstre o endereço deste valor na memória.
// - Crie um struct "pessoa"
// - Crie uma função chamada mudeMe que tenha *pessoa como parâmetro.
//Essa função deve mudar um valor armazenado no endereço *pessoa.
// - Dica: a maneira "correta" para fazer dereference de um valor em um struct
//seria (*valor).campo
// - Crie um valor do tipo pessoa;
// - Use a função mudeMe e demonstre o resultado.
package main

import "fmt"

type Person struct {
	Name    string
	Surname string
	Age     int
	Country string
}

func mudeMe(p *Person) string {
	y := p.Surname
	return y
}

func main() {

	x := 10
	fmt.Println(&x)

	Alexandro := Person{"Alexandro", "Dias", 39, "Estados Unidos"}
	fmt.Println(Alexandro)
	fmt.Println(mudeMe(&Alexandro))

}
