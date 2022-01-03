package main

import (
	"fmt"
	// Dentro de math há outros pacotes, importamos apenas aquele que iremos usar
	"math/rand"
)

func main() {
	fmt.Println("Meu número favorito é", rand.Intn(10))
}