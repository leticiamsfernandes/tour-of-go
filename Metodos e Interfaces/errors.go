package main

import (
	"fmt"
	"time"
)

// Programas Go expressam estados de erro com valores error.
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

// O tipo error é uma interface embutida similar a fmt.Stringer.
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	// Funções frequentemente retornam um valor error e a chamada do código deve lidar com erros testando se o erro é igual a nil.
	// Um error nil indica sucesso; um error não-nil indica fracasso.
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
