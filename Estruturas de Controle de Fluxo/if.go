package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	//A expressão não precisa ser cercada de ( ) mas os chaves { } são obrigatórios.
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
