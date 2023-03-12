// - Crie uma função que retorna uma função.
// - Atribua a função retornada a uma variável.
// - Chame a função retornada.

package main

import "fmt"

func main() {

	x := funcao1()
	x()

}

func funcao1() func() {
	return func() {
		fmt.Println("Uma função que retorna uma funcao")
	}
}
