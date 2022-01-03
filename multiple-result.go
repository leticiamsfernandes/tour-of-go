package main

import "fmt"

//Especifica-se o tipo do retorno e quantidade de dados retornada, na ordem
func swap(x, y string) (string, string) {
	//Uma função pode retornar qualquer número de resultados.
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
