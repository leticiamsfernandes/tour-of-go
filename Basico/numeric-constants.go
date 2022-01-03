package main

import "fmt"

//Constantes numéricas também são valores de alta-precisão.
//Uma constante sem tipo pega o tipo necessário pelo seu contexto.

const (
	// Cria um número enorme deslocando 1 bit para a esquerda 100 casas.
	// Em outras palavras, o número binário que é 1 seguido por 100 zeros.
	Big = 1 << 100
	// Desloca para a direita novamente 99 casas, então termina com 1 << 1 ou 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
