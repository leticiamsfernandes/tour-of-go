package main

import (
	"fmt"
	"math"
)

// Um tipo de interface é definida por um conjunto de métodos.
type Abser interface {
	// Um valor de tipo de interface pode conter qualquer valor que implementa esses métodos.
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // o MyFloat implementa Abser
	a = &v // o *Vertex implementa Abser

	// Na linha seguinte, v é do tipo Vertex (não *Vertex)
	// e NÃO implementa Abser.
	//a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
