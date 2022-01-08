package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// Uma goroutine é um segmento leve e gerenciado pelo runtime de Go.
	go say("world")
	say("hello")
}

// Goroutines executam no mesmo espaço de endereço, para que o acesso à memória compartilhada seja sincronizada.
// O pacote sync fornece as primitivas úteis, embora você não vá precisar muito deles em Go, pois há outras primitivas.
