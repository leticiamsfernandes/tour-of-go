package main

import (
	"fmt"
	"math"
)
// Há duas razões para usar um receptor de ponteiro.
// O primeiro é para que o método possa modificar o valor que os seus pontos de receptor referenciam.
//A segunda é para evitar copiar o valor de cada chamada de método. Isto pode ser mais eficaz se o receptor for uma estrutura grande, por exemplo

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Em geral, todos os métodos em um determinado tipo devem ter o valor ou ponteiro receptores, mas não uma mistura de ambos. 

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
