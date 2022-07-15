package main

import "fmt"

func main() {
	// variavel que salva dentro dela não necessariamente um valor e sim o endereço de memoria de alguma coisa

	var variavel1 int = 10
	var variavel2 int = variavel1
	fmt.Println(variavel1, variavel2)
	//atribuir valor por copia, somente mudará até determinada linha de código, neste caso, somente a variavel1 teve alteração, pois a v2 é como se fosse uma cópia
	variavel1++
	fmt.Println(variavel1, variavel2)

	//PONTEIRO É UMA REFERENCIA DE MEMORIA

	var variavel3 int = 100
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, *ponteiro, ponteiro) // se quisermos ver o valor atraves do ponteiro, colocamos um asterisco: chama-se de desreferenciação

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro)

	//memoria do seu pc tem varias gavetas, qnd vc cria uma variavel ela fica nessas gavetas, qnd vc cria um ponteiro, vc n passa oq ta dentro da gaveta e sim o endereço dela, ex: 'Tá nessa gaveta aqui'

}
