package main

import "fmt"

type Vertex struct {
	X, Y int
}

// A struct literal indica um valor struct recém-alocado, ao enumerar os valores de seus campos.
// Você pode listar apenas um subconjunto de campos usando o Name: sintaxe. (E a ordem dos campos nomeados é irrelevante.)
var (
	v1 = Vertex{1, 2}  // possui o tipo Vertex
	v2 = Vertex{X: 1}  // Y:0 é implícito
	v3 = Vertex{}      // X:0 and Y:0
	// O prefixo especial & constrói um ponteiro para uma struct litera
	p  = &Vertex{1, 2} // possui o tipo *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}