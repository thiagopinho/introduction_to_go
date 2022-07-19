package main

import (
	"fmt"
)

func main() {
	// Duas maneiras de declarar 1. deixando o tipo inplicito ou explicito

	var variavel1 int = 1 // Declaração explícita, estamos colocando o tipo diretamente

	tipoOculto := 3 // Declaração implícita, a variável declara seu tipo com base no valor recebido (Inferência de tipos)

	fmt.Println("Valor normal das variáveis", variavel1, tipoOculto)

	// const oi := "const"
	const constant string = "alo"

	// fmt.Print(constant)

	variavel1, tipoOculto = tipoOculto, variavel1
	fmt.Println("Inversão das variáveis", variavel1+tipoOculto)
	// Não precisa de uma variável terceira para inverter os valores, o Go consegue fazer isso.
}

// O go não deixa você declarar variável e não usar, o mesmo vale para pacote, isso pode parecer chato mas em projeto grande isso acaba sendo útil, pois normalmente podemos esquecer uma variável ou outra no código.
