package main

import "fmt"

func main() {

	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restodaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restodaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	//nao poderiamos, por exemplo usar int16 e somar com um int32
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	// Fim aritméticos

	// Atribuiçao

	var variavel string = "String"
	variavel2 := "String2"

	fmt.Println(variavel, variavel2)

	// Fim dos operadores de atribuição

	// Operadores relacionais retornam bool
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	// FIM DOS RELACIONAIS

	//operadores logicos
	fmt.Println("----------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro) //sinal de negação, inverter o valor de uma variável booleana

	//FIM DOS OPERADORES LOGICOS

	//OPERADORES UNARIOS
	numero := 10
	numero++
	numero += 15 //numero + numero + 15
	fmt.Println(numero)

	numero--
	numero -= 20
	fmt.Println(numero)

	numero *= 3
	numero /= 10
	numero %= 3

	fmt.Println(numero)

	//FIM OPERADOR UNARIO

	//nao existe ternario em Go, uma forma de contornar seria o conhecido if-else

	var texto string
	if numero > 5 {
		texto = "maior que 5"
	} else {
		texto = "menor que 5"
	}

	fmt.Println(texto)

	// A premissa do Go é que existe apenas uma maneira de fazer, a exceção nessa regra seria a declaração de variáveis
}
