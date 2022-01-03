package main

import "fmt"

//A declaração var pode incluir inicializadores, uma por variável.
var i, j int = 1, 2

func main() {
	//Se um inicializador está presente, o tipo pode ser omitido; a variável terá o tipo do inicializador.
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}