// - Crie e use um struct anônimo.
// - Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.

package main

import "fmt"

func main() {

	structAnonino := struct {
		nome      string
		sabor     string
		ondeTem   []string
		vaiBemCom map[string]string
	}{
		nome:    "Donuts",
		sabor:   "chocolate",
		ondeTem: []string{"Nos EUA", "na Europa"},
		vaiBemCom: map[string]string{
			"Donuts": "Vai bem com suco",
		},
	}
	fmt.Println(structAnonino)
}
