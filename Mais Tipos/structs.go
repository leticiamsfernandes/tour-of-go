package main

import "fmt"

// A struct é uma coleção de campos.
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}