package main

import "fmt"

// O valor zero de uma slice é nil.
// Uma slice nula tem o comprimento e a capacidade igual 0.
func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}