package main

import "fmt"

// Quando dois ou mais parâmetros nomeados consecutivos de uma função compartilham o mesmo tipo, você pode omitir o tipo de todos menos o último.
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}