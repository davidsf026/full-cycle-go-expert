// Não funciona, porque ele por default procura na pasta src da $GOROOT

package main

import (
	"fmt"
	"matematica"
)

func main() {
	s := matematica.soma(10, 20)
	fmt.Println(s)
}
