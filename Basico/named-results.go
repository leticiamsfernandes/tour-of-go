package main

import "fmt"

func split(sum int) (y, x int) /*Valores de retorno podem ser nomeados e devem ser usados para documentar o significado dos valores de retorno*/ {
	x = sum * 4 / 9
	y = sum - x
	/*A declaração "return" sem argumentos retorna os valores atuais dos dois resultados, e é conhecido como retorno limpo. 
	Esse tipo de retorno deve ser usado apenas em funções curtas*/
	return
}

func main() {
	fmt.Println(split(17))
}