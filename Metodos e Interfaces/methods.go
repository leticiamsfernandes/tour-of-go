package main

import (
	"fmt"
	"math"
)

// Go não tem classes. No entanto, você pode definir métodos em tipos.
type Vertex struct {
	X, Y float64
}

// O método é uma função com um argumento receptor especial.
// O receptor aparece em sua lista de argumentos entre a própria palavra-chave func e o nome do método.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}

	// Neste exemplo, o método Abs tem um receptor do tipo Vertex chamado v.
	fmt.Println(v.Abs())
}
