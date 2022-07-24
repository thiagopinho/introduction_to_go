package main

import (
	"fmt"
	"sync"
	"time"
)

// Concorrência != Paralelismo

// Paralelismo = acontece quando tem duas ou mais tarefas sendo executadas ao mesmo tempo, precisa ter mais de um núcleo no processaodor, pois ele irá distribuir as tarefas entre os núcleos e elas serão executadas ao mesmo tempo de fato

// concorrente => não necessariamente estão executando ao mesmo tempo, podem também, caso você esteja com um processador de mais um núcleo, mas não necessariamente, se vc tiver um processador com apenas um núcleo, vc tbm pode aplicar concorrência, a diferença maior seria que uma tarefa não esperaria outra acabar, elas iriam se revezando.

func main() {

	var waitGroup sync.WaitGroup

	waitGroup.Add(2) //quantidade de go routines, grupo de espera, precisamos espeficicar quantas goroutines rodarão ao mesmo tempo, ele não faz esse cálculo sozinho

	go func() {
		escrever("Olá mundo!")
		waitGroup.Done() // tira um do contador (-1)
	}()

	go func() {
		escrever("Programando em Go")
		waitGroup.Done() // tira um do contador (-1)
	}()

	waitGroup.Wait() // 0 (quando chegar a 0, as próximas etapas do Programa podem prosseguir, sem essa cláusula, ele não saberia quando parar)

	// Quando colocamos a palavra Go: Execute a função mas não espera ela terminar de executar para seguir o programa

	go escrever("Olá mundo") //goroutine
	escrever("Programando em GO")
}

func escrever(texto string) {

	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}

}
