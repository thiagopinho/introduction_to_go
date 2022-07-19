package main

import "fmt"

//funcao recursiva não é recomendada em casos com muitas execuções, pois pode demorar pois uma função dependeria de outra e ficaria muito pesado.

func fibonacci(posicao uint) uint {
	//funcoes recursivas precisam de uma parada pois pode ficar em execução inifinita e gerar um estouro de pilha (stackoverflow) = cada execução é uma pilha, se tiver muitas, ela terá uma lista, se vc nunca parar, essa lista será gigante e seu sistema não comportará.

	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)

}

func main() {
	fmt.Println("Funções Recursivas")
	posicao := uint(10)

	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}
	fmt.Println(fibonacci(posicao))
	// 1 1 2 3 5 8 13 == fibonacci

}
