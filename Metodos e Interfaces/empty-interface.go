package main

import "fmt"

// Uma interface vazia pode conter valores de qualquer tipo.

// Interfaces vazias são usadas pelo código que lida com valores de tipo desconhecido. 
// Por exemplo, fmt.Print leva qualquer número de argumentos do tipo interface{}.

func main() {
	// O tipo de interface que especifica zero métodos é conhecido como interface vazia.
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
