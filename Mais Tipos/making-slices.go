package main

import "fmt"

// As slices podem ser criadas com a função make padrão; isto é como você criar matrizes dinamicamente.
// O função make aloca uma matriz zerada e retorna uma slice que se refere a essa matriz e para especificar a capacidade, passe o terceiro argumento no make.
func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}