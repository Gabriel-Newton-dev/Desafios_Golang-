// - Utilize a declaração defer de maneira que demonstre que sua execução
// só ocorre ao final do contexto ao qual ela pertence.

package main

import "fmt"

func main() {

	x := 10
	y := 15
	soma := x + y
	multiplicacao := x * y

	defer fmt.Println(soma)
	fmt.Println(multiplicacao)

}
