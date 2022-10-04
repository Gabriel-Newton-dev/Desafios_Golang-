package main

import "fmt"

var DigitarSenha int

func main() {

	fmt.Print("Digite uma senha com 4 números:")
	fmt.Scan(&DigitarSenha)

	if DigitarSenha == 1234 {
		fmt.Println("Acesso liberado")
	} else {
		fmt.Println("Senha incorreta, acesso negado")
	}

}
