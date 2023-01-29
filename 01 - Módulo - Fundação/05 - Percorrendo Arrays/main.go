package main

import (
	"fmt"
	"strconv"
)

func main() {
	start := 1
	things := []string{
		"maleta",
		"bolso",
		"boleto",
		"vaso",
		"casa"}

	fmt.Println(things)

	fmt.Println("===============================")

	// FOR TRADICIONAL
	for i := 0; i < 5; i++ {
		fmt.Println("Valor: " + strconv.Itoa(i+start))
	}

	fmt.Println("===============================")

	for i := 0; i < len(things); i++ {
		fmt.Println(things[i])
	}

	fmt.Println("===============================")

	// FOR EACH PEGANDO APENAS VALOR
	for i := range things {
		fmt.Println(things[i])
	}

	fmt.Println("===============================")

	// FOR EACH PEGANDO INDICE E VALOR
	for i, v := range things {
		fmt.Printf("O índice %d tem valor %s\n", i, v)                 // JEITO MALUCO
		fmt.Println("O índice " + strconv.Itoa(i) + " tem valor " + v) // JEITO SÃO (NA MINHA OPINIÃO RSRSRSRSR)
		fmt.Println("")
	}
}
