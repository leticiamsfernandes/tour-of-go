package main

import "fmt"

//Constantes são declaradas como variáveis, mas com a palavra-chave const.
//Podem ser seqüências de caracteres, booleanos, ou valores numéricos.
//Não podem ser declaradas usando a sintaxe :=.

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
