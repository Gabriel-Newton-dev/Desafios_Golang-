// Callback: passe uma função como argumento a outra função.

package main

import "fmt"

func main() {
	recebeCallback(callback)
}

func callback() {
	fmt.Println("Essa será a nossa função para realizar o callback ")
}

func recebeCallback(x func()) {
	fmt.Println("Irá receber o call back")
	x()
}
