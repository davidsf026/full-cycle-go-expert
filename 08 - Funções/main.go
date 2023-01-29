package main

import "fmt"

func main() {
	fmt.Println(sum(2, 3))
	count(10)

	result, fiftyFour := sumFiftyFour(22, 32)

	println(result)
	println(fiftyFour)
}

// Função Com Retorno de Inteiro
func sum(a, b int) int {
	return a + b
}

// Função Void (sem retorno)
func count(a int) {
	for i := 0; i < a; i++ {
		fmt.Println(i + 1)
	}
}

// Função que retorna dois valores
func sumFiftyFour(a, b int) (int, string) {
	if a+b == 54 {
		return (a + b), "LET'S GO TO FIFTY FOUR!!!"
	} else {
		return (a + b), "boring..."
	}
}
