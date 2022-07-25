package main

import "fmt"

func main() {

	canal := make(chan string, 3) //especifica a capacidade do canal, canal com buffer só vai bloquear quando atingir a capacidade máxima
	canal <- "Olá mundo!"
	canal <- "Programando em GO"
	canal <- "Mensagem 3"
	//canal <- "Terceiro valor" // dá deadlock se tiver com 2
	mensagem := <-canal
	mensagem2 := <-canal
	mensagem3 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
	fmt.Println(mensagem3)
}
