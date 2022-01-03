package main

import "fmt"

// O valor resultante do append é uma slice contendo todos os elementos da slice original mais os valores providos.
// Se a matriz anterior de s for muito pequena para caber todos os valores uma matriz gigante será alocada. A slice retornada apontará para a nova matriz alocada.
func main() {
	var s []int
	printSlice(s)

	// append funciona em slices nulas.
	s = append(s, 0)
	printSlice(s)

	// A slice cresce conforme a necessidade.
	s = append(s, 1)
	printSlice(s)

	// Podemos adicionar mais de um elemento por vez.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
