package main

import "fmt"

// A type assertation fornece acesso ao valor concreto subjacente de um valor de interface.
func main() {
	var i interface{} = "hello"

	// Esta declaração afirma que o valor de interface i detém o tipo concreto string e atribui o valor de T subjacente à variável s.
	s := i.(string)
	fmt.Println(s)

	// Se i detém T, então s será o valor subjacente e ok será true.
	// Se não, ok será falso e s será o valor zero do tipo T, e sem ocorrer erro pânico.
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// Se i não detém uma T, a declaração irá desencadear um erro pânico.
	f = i.(float64) // panic
	fmt.Println(f)
}
