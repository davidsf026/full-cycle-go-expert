// Arrays não podem ter a sua capacidade de índices alteradas em tempo de execução. Já myFruitSlices, sim!

package main

import "fmt"

func main() {
	myFruitSlice := []string{ // SLICES NÃO DECLARAM O VALOR ENTRE []
		"banana",
		"maçã",
		"abacate",
		"tomate",
		"abóbora"}

	fmt.Println("ID\tLEN\tCAP\tDATA")

	// 1. MOSTRANDO TODOS OS INDICES
	fmt.Printf("1\t%d\t%d\t%s\n", len(myFruitSlice), cap(myFruitSlice), myFruitSlice)

	// 2. MOSTRANDO OS INDICES ENTRE O COMEÇO DO SLICE ATÉ O TRÊS
	fmt.Printf("2\t%d\t%d\t%s\n", len(myFruitSlice), cap(myFruitSlice), myFruitSlice[:3])

	// 3. MOSTRANDO OS INDICES ENTRE O INDICE 3 E O FINAL DO SLICE
	fmt.Printf("3\t%d\t%d\t%s\n", len(myFruitSlice), cap(myFruitSlice), myFruitSlice[3:])

	// 4. MOSTRANDO OS INDICES ENTRE O INDICE 2 E 4
	fmt.Printf("4\t%d\t%d\t%s\n", len(myFruitSlice), cap(myFruitSlice), myFruitSlice[2:4])

	///////////////////////////////////////////////////

	myFruitSlice = append(myFruitSlice, "kiwi")

	// 5. MOSTRANDO SLICE APÓS APPEND
	fmt.Printf("5\t%d\t%d\t%s\n", len(myFruitSlice), cap(myFruitSlice), myFruitSlice) // REPARE QUE A CAPACIDADE DOBROU MESMO EU TENDO ADICIONADO APENAS UM, ESSA É UMA DESVANTAGEM

}
