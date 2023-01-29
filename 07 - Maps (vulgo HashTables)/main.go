package main

import "fmt"

func main() {
	// Maps são listas de chave-valor ou, em outras palavras, um array com índice personalizado.

	// Como declarar um map "[tipo-da-chave]tipo-do-valor":
	salarios := map[string]int{
		"David":   30000,
		"João":    1000,
		"Rafaela": 8000,
	}

	// Mostrando tudo o que tem no map
	fmt.Println(salarios)

	// Apagando um item
	delete(salarios, "Rafaela")
	fmt.Println(salarios)

	// Retonar o valor de um map pela chave
	fmt.Printf("O salário do David é: %d\n", salarios["David"])

	// Percorrendo um map
	for i := range salarios {
		fmt.Println(i)
	}

	for _, v := range salarios {
		fmt.Println(v)
	}

	for i, v := range salarios {
		fmt.Println(i, v)
	}
}
