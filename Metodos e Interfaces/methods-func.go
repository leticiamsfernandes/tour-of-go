package main

import (
	"fmt"
	"math"
)

// Um método é apenas uma função com um argumento receptor.
type Vertex struct {
	X, Y float64
}

// Aqui está Abs escrito como uma função regular, sem qualquer alteração na funcionalidade.
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
