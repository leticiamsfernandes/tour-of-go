package main

import "fmt"

//Go tem apenas uma estrutura de laço, o for.
func main() {
	sum := 0

	//Não há parênteses e a chave é sempre necessária.
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}