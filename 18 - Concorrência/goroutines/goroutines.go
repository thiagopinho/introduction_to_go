package main

import (
	"fmt"
	"time"
)

// Concorrência != Paralelismo

// Paralelismo = acontece quando tem duas ou mais tarefas sendo executadas ao mesmo tempo, precisa ter mais de um núcleo no processaodor, pois ele irá distribuir as tarefas entre os núcleos e elas serão executadas ao mesmo tempo de fato

// concorrente => não necessariamente estão executando ao mesmo tempo, podem também, caso você esteja com um processador de mais um núcleo, mas não necessariamente, se vc tiver um processador com apenas um núcleo, vc tbm pode aplicar concorrência, a diferença maior seria que uma tarefa não esperaria outra acabar, elas iriam se revezando.

func main() {
	// Quando colocamos a palavra Go: Execute a função mas não espera ela terminar de executar para seguir o programa

	go escrever("Olá mundo") //goroutine
	escrever("Programando em GO")
}

func escrever(texto string) {

	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}

}
