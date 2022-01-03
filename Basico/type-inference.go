package main

import "fmt"

//Ao declarar uma variável, sem especificar o seu tipo, o tipo da variável é inferida a partir do valor do lado direito.
func main() {
	v := 1.456
	fmt.Printf("v is of type %T\n", v)
}