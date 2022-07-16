package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	numero := 10

	// No Go nao precisamos colocar parenteses, somente se precisa colocar em uma ordem de prescência como uma soma etc.

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	//inicializa uma variavel nessa condição = if init
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que 0")
	} else {
		fmt.Println("Número não é menor que 0")
	}

	//fmt.Println(outroNumero) // if init = limita ao escopo do if, não podemos nesse caso acessar a variável "outroNumero"
}
