package main

func main() {
	println(infinitySum(2, 3, 4, 5, 6, 6, 7, 7, 8, 9, 5))
}

// Função Com Retorno de Inteiro
func infinitySum(values ...int) int {
	var result int

	for _, i := range values {
		result += i
	}

	return result
}
