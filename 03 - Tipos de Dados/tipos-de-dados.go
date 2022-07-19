package main

import (
	"errors"
	"fmt"
)

func main() {
	//Temos dois tipos de números, inteiros e reais, eles se subdividem entre várias categorias

	// No Go são 4 tipos de inteiros

	// int8, int16, int32, int64

	var inteiro8 int8 = 100
	var inteiro16 int16 = 10000
	var inteiro32 int64 = 1000000000000000000
	var inteiro64 int64 = 1000000000000000000
	// o int sozinho, quando nao especificamos, usa a arquitetura do computador como base, exemplo: Se o seu computador for 32 bits, ele criará um int32, caso seja 64 bits, int64

	fmt.Println("INTEIROS:", inteiro8, inteiro16, inteiro32, inteiro64)

	//alias = apelido

	//RUNE = INT32, USADO PARA SE REFERIR A NUMEROS QUE REPRESENTAM CARACTERES/tabela ascii

	//float32, float64

	var real32 float32 = 123.123123
	var real64 float64 = 123.12312384108
	// * Declarar somente como float dará erro, precisamos declarar se será 32 ou 64

	fmt.Println("REAIS:", real32, real64)

	char := 'T'

	fmt.Println(char)
	// * Printa o valor do número na tabela ASCII, o Go não possui o tipo char.

	// ^ Valor zero, todo tipo de dado tem o seu valor zero, exemplo: String = "", int = 0, erro = nil, bool = false

	var texto string
	fmt.Println(texto)

	var boolFalse bool = false // (caso declare bool como false, simplesmente declare var boolFalse bool, sem passar o valor)
	var boolTrue bool = true

	fmt.Println(boolFalse, boolTrue)

	var erro error = errors.New("Erro interno") // serve como valor zero para interface, erro, ponteiros
	//errors é um pacote nativo do Go, caso queira personalizar a mensagem do erro, terá que utilizar este pacote.

	fmt.Println(erro)
}
