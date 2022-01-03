package main

// O pacote io especifica a interface io.Reader, que representa o fim da leitura de um fluxo de dados.
import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		// A interface io.Reader tem um método Read.
		/* Read popula uma slice de bytes passados com dados e retorna o número de bytes populados e um valor de erro. 
		Este retorna um erro io.EOF quando o fluxo termina. */
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
