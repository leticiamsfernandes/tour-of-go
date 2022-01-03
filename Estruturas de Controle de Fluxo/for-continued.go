package main

import "fmt"

func main() {
	sum := 1
	//A instrução inicial e a pós-instrução são opcionais.
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}