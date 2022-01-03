package main

import "fmt"

/* Uma matriz tem um tamanho fixo. Uma slice, por outro lado, é dinamicamente redimensionada, uma visão flexível dos elementos de uma matriz. 
Na prática, as slices são muito mais comuns do que as matrizes. */
// O tipo []T é uma slice com elementos do tipo T.
func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// Uma slice é fomada pela especificação de dois indices, um limite menor e maior, separados por dois pontos.
	//Este seleciona um intervalo meio-aberto que inclui o primeiro elemento, mas exclui o último.
	var s []int = primes[1:4]
	fmt.Println(s)
}