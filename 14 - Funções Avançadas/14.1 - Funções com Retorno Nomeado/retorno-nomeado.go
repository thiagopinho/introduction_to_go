package main

import "fmt"

func calculosMatematico(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return

	// return soma, subtracao (a ideia do retorno nomeado é não precisar declarar todas variaveis no retorno, apenas o "RETURN")
}

func main() {
	soma, subtracao := calculosMatematico(10, 20)
	fmt.Println(soma, subtracao)
}

// Em alguns casos vale a pena usar o retorno nomeado pois fica mais legível
