package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	// Campos Struct são acessados ​​através de um ponto.
	v.X = 4
	fmt.Println(v.X)
}
