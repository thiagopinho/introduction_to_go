package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {

		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {
		//A função select permite que não fiquemos esperando a execução de uma função a espera de outra, no caso a execução do canal 1 é e deveria ser feita muito mais vezes do que a do canal 2, e não deveria ser dependente dele.
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)

		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}

	}

}
