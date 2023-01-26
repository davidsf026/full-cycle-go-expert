package main

import "fmt"

// Se uma variável nasce string ela vai morrer string, não dá mudar o tipo dela.

const a = "Hello, World!" // É um valor que não pode ser mudado, é constante
// a = "Hello, Moon!" // Não dá pra fazer isso.

// FORMA TRADICIONAL DE DECLARAR VAR
var b bool // Inicializado por default como false

// FORMA EM ARRAY DE DECLARAR VAR
var (
	c int            // Inicializado por default como 0
	d string         // Inicializado por default como ""
	e float64        // Inicializado por default como +0.000000e+000
	f bool    = true // Forcei a inicialização como true
)

// CRIANDO TIPO PERSONALIZADO
type ID int

var g ID // Eis uma variável do tipo ID que por debaixo dos panos é um int

func main() {
	// FORMA RESUMIDA DE DECLARAR VAR
	h := "Isso é uma String e o Go sabe disso." // Go sabe que é string e portanto o tipo de g será string. Aparentemente não ser para var global.

	b = true
	println(a)
	fmt.Printf("O tipo da var a é: %T", a)
	println(b)
	println(c)
	println(d)
	println(e)
	println(g)
	fmt.Printf("O tipo da var g é: %T", g)
	println(h)

	// DEFININDO UM ARRAY
	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 3

	println(meuArray[0])
	println(meuArray[1])
	println(meuArray[2])

	println(len(meuArray))
	println(meuArray[len(meuArray)-2])
}
