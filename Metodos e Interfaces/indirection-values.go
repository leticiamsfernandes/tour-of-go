package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	// Funções que recebem um argumento de valor devem ter um valor desse tipo específico, não pode ser a referência do ponteiro.
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}

	// O método que chama p.Abs() é interpretado como (*p).Abs().
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
