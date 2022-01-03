package main

import "fmt"

// O tipo vem depois do nome da variável
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}