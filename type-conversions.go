package main

import (
	"fmt"
	"math"
)

//A expressão T(v) converte o valor v para o tipo T.

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	//Em Go atribuição entre os itens de tipo diferente requer uma conversão explícita.
	var z uint = uint(f)
	fmt.Println(x, y, z)
}