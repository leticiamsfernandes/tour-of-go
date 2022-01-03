package main

import "fmt"

// Um tipo implementa uma interface através da implementação dos métodos.
// Interfaces implícitas dissociam-se da definição de uma interface de sua implementação, que pode então aparecer em qualquer pacote sem pré-arranjamento.
type I interface {
	M()
}

type T struct {
	S string
}

// Este método significa que o tipo T implementa a interface I,
// mas não precisamos declarar explicitamente que isso acontece.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
