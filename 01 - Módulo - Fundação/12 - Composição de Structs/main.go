package main

type Endereco struct {
	Casa   string
	Rua    string
	Bairro string
	Cidade string
	Pais   string
}

// Composição
type FuncionarioComposicao struct {
	Nome    string
	Idade   int
	Salario int
	Endereco
}

// Juntando duas structs por instanciação de uma dentro da outra
type Funcionario struct {
	Nome    string
	Idade   int
	Salario int
	Adress  Endereco
}

func main() {

}
