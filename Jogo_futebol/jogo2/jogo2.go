// Caso 2: Várias partidas
// A = [2,0,2,3,2]
// B = [1,1,2,1,2]
// len(A) = len(B)

package main

import "fmt"

func main() {

	A := []int{2, 0, 2, 3, 2}
	B := []int{1, 1, 2, 1, 2}

	if A[0:5] > B[0:5] {
		fmt.Println("Time A ganhou")
	}

}
