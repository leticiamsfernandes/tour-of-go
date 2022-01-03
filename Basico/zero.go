package main

import "fmt"

//Variáveis declaradas sem um valor inicial explicitado darão seu valor zero.
func main() {
	var i int
	var f float64
	var b bool
	var s string
	/* 
		0 para tipos numéricos,
		false para tipos boleanos, e
		"" (string vazia) para strings.
	*/
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
