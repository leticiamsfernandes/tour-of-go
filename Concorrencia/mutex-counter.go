package main

import (
	"fmt"
	"sync"
	"time"
)

/* Ter certeza que uma única goroutine pode acessar uma variável de cada vez para evitar conflitos é chamado exclusão mútua, 
e o nome convencional para a estrutura de dados que a fornece é mutex. */

/* Podemos definir um bloco de código a ser executado em exclusão mútua pelo que o rodeia com uma chamada para Lock e Unlock como mostrado no método Inc.
Nós também podemos usar defer para garantir que o mutex será desbloqueado como no método Value. */

//  SafeCounter é seguro para uso concorrente.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc incrementa o contador para a chave fornecida.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock para que apenas uma goroutine por vez possa acessar o mapa c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value retorna o valor atual do contador para a chave fornecida.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock portanto, apenas uma goroutine por vez pode acessar o mapa c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
