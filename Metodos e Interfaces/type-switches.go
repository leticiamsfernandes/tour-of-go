package main

import "fmt"

// Um type switch é uma construção que permite diversas type assertations em série.
/* Um type switch é como uma instrução switch regular, mas os cases em um type switch especificam os tipos (e não valores), 
e esses valores são comparados com o tipo dos valores determinados da interface informada. */
func do(i interface{}) {
	// A declaração em um _type switch_ tem a mesma sintaxe como uma afirmação de tipo de i.(T), mas o tipo específico t é substituído com a palavra-chave type.
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
