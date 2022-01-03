package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // aponta para i
	fmt.Println(*p) // lê i através do ponteiro
	*p = 21         // seta i através do ponteiro
	fmt.Println(i)  // imprime o novo valor de i

	p = &j         // aponta para j
	*p = *p / 37   // divide j através do ponteiro
	fmt.Println(j) // imprime o novo valor de j
}