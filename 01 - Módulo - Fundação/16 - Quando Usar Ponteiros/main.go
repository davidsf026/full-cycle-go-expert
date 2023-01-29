package main

func alterarValorNaMemoria(a *int) {
	*a = 54
}

func main() {
	varA := 10
	println(varA)

	alterarValorNaMemoria(&varA)

	println(varA)
}
