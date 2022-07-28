//fila de tarefas e vc tem varios workers pegando itens dessa fila de maneira independente

package main

import "fmt"

func main() {
	tarefas := make(chan int, 45)    // sequencia de numeros
	resultados := make(chan int, 45) // armazena a medida do calculo

	go worker(tarefas, resultados)
	go worker(tarefas, resultados) // colocando mais um worker o código gera mais rápido
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	//chegará um ponto que nn importa quantas go routines vc coloca pq o resultado nao vai ser tao diferente assim, pois é limitado ao processador da máquina

	for i := 0; i < 45; i++ {
		tarefas <- i
	}

	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}

	print(worker)

}

//tarefas canal que só envia <-
//resultados canal que só recebe

func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
