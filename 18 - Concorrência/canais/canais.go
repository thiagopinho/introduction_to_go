package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Olá mundo", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa!")

	//canal <- valor = mandando valor pro canal
	//<- canal = recebendo um valor, esperando que o canal receba um valor
}

func escrever(texto string, canal chan string) {
	time.Sleep(time.Second)
	for i := 0; i < 5; i++ {
		canal <- texto //canal recebe o texto
		time.Sleep(time.Second)
	}

	close(canal) //depois que terminar o loop de 5 vezes, vamos fechar o canal
}

//dead lock = Você não tem mais nenhum lugar que tá enviando dado pro teu canal, mas seu canal ainda espera receber um dado, deadlock não é pego em compilação, somente execução.(tomar cuidado ao subir em produção)

// Canal tem duas operações: Enviar e receber um dado

// Essas operações bloqueiam a execução do programa
