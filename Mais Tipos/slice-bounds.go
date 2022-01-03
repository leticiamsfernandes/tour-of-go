package main

import "fmt"

// Ao "fatiar", você pode omitir as posições de limite altas ou baixas para usar seus padrões em seu lugar.
// O padrão é igual a zero para o limite baixo e o comprimento da slice para o limite alto.
func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}