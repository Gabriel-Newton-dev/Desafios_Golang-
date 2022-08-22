// Escreva  um  programa  que  verifique  a  validade  de  uma  senha  fornecida  pelo  usuário.
// A  senha  válida  é  o  número  1234. Devem  ser impressas  as  seguintes mensagens:
// ACESSO PERMITIDO caso a senha seja válida.
// ACESSO NEGADO caso a senha seja inválida.

package main

import "fmt"

var digiteUmaSenha int

func main() {

	senhaValida := 1234

	fmt.Print("Digite uma senha de 4 números por favor: ")
	fmt.Scan(&digiteUmaSenha)

	if senhaValida == digiteUmaSenha {
		fmt.Println("ACESSO PERMITIDO")
	} else {
		fmt.Println("ACESSO NEGADO ")
	}

}
