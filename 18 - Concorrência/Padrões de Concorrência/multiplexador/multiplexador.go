package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	canal := multiplexar(escrever("Olá mundo"), escrever("Programando em GO"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexar(canalDeEntrada1, canalDeEntrada2 <-chan string) <-chan string {
	//pegar um ou mais canais e implementar em um só
	canalDeSaida := make(chan string)

	//encaminhamento de mensagem
	go func() {
		for {
			select {
			case mensagem := <-canalDeEntrada1:
				canalDeSaida <- mensagem

			case mensagem := <-canalDeEntrada2:
				canalDeSaida <- mensagem
			}
		}
	}()

	return canalDeSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)
	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000))) //calcula quanto tempo demorou pra executar cada operação
		}
	}()

	return canal
}
