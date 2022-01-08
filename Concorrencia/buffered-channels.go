package main

import "fmt"

func main() {
	// Os canais podem ser bufferizados. Fornecendo o tamanho do buffer como o segundo argumento para make para inicializar um canal bufferizado.
	// Envia para um bloco de canais bufferizados apenas quando o buffer está cheio. Recebe bloco quando o buffer está vazio.
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
