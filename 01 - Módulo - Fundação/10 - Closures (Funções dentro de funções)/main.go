package main

func main() {

	myClosure := func() int {
		return infinitySum(1)
	}()

	println(myClosure)

}

// Função Com Retorno de Inteiro
func infinitySum(values ...int) int {
	var result int

	for _, i := range values {
		result += i
	}

	return result
}
