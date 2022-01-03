package main

import "fmt"

func main() {
	m := make(map[string]int)

	// Inserir ou atualizar um elemento no map m
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	// Recuperar um elemento
	fmt.Println("The value:", m["Answer"])

	// Excluir um elemento
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// Teste que uma chave está presente com dois valores.
	// Se "Answer" está em m, ok é true. Se não, ok é false.
	// Se "Answer" não está no map então v tem valor zero para o elemento do tipo do map.
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
