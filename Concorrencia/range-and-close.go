package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// Um remetente pode close um canal para indicar que os valores não serão mais enviados. 
	// Apenas o remetente deve fechar um canal, nunca o receptor. O envio em um canal fechado irá causar um pânico.
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// O laço for i := range c recebe valores do canal repetidamente até que seja fechado.
	for i := range c {
		fmt.Println(i)
	}
}

// Canais não são como arquivos, você geralmente não precisa fechá-los. 
// O encerramento só é necessário quando o receptor precisa saber que não há mais valores chegando, como para terminar um laço range.
