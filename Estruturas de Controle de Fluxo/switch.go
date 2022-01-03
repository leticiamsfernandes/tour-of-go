package main

import (
	"fmt"
	"runtime"
)

//O Go apenas executa o caso selecionado, não todos os casos que seguem.
//A declaração break que é necessária no final de cada caso naquelas linguagens é fornecido automaticamente no Go.
//Outra diferença importante é que os cases dos switchs do Go não precisam ser constantes, e os valores envolvidos não precisam ser inteiros.
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}