package main

// Como declarar uma classe, repare que você está criando um tipo personalizado de uma struct
type Funcionario struct {
	Nome    string
	Idade   int
	Salario int
}

func main() {
	// Instanciar um objeto da classe
	David := Funcionario{
		Nome:    "David",
		Idade:   19,
		Salario: 30000,
	}

	David.Salario = 60000

	println(David.Nome, David.Salario)
}
