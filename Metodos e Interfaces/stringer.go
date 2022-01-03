package main

import "fmt"

// Uma das interfaces mais ubíqua é a Stringer definida pelo pacote fmt.
// A Stringer é um tipo que pode descrever-se como uma string. O pacote fmt (e muitos outros) olham para essa interface para mostrar valores.

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}