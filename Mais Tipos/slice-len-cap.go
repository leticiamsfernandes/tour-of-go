package main

import "fmt"

// Uma slice tem tanto um tamanho quanto uma capacidade.
// O comprimento de uma slice é o número de elementos que ela contém.
// A capacidade de uma slice é o número de elementos na matriz subjacente, contando a partir do primeiro elemento na slice.
// O comprimento e a capacidade de uma slice S podem ser obtidos utilizando as expressões len(s) e cap(s).


func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Divide a fatia para obter comprimento zero.
	s = s[:0]
	printSlice(s)

	// Aumenta seu comprimento.
	s = s[:4]
	printSlice(s)

	// Apaga seus primeiros dois valores.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
