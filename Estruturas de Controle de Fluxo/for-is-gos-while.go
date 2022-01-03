package main

import "fmt"

func main() {
	sum := 1
	//Pode tirar as vírgulas: o while do C é escrito com for em Go.
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}