package main

import (
	"fmt"
	"math"
)

func main() {
	// Em Go, um nome é exportado se ele começa com uma letra maiúscula. 
	// Ao importar um pacote, você pode referenciar apenas seus nomes exportados. Todos os nomes "não exportados" não são acessíveis de fora do pacote.
	fmt.Println(math.Pi)
}
