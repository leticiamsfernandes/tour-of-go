package main

import "fmt"

// Canais são um conduto tipado através do qual você pode enviar e receber valores com o operador de canal, <-.
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// Os dados fluem na direção da seta.
	c <- sum // coloca a soma em c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	// Como maps e slices, os canais devem ser criados antes de se usar.
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // recebe de c

	fmt.Println(x, y, x+y)
}

// Por padrão, enviam e recebem bloco até o outro lado estar pronto. Isso permite que goroutines sincronizem sem bloqueios explícitos ou variáveis ​​de condição.