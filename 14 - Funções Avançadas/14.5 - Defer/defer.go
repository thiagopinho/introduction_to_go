package main

import "fmt"

//defer = adia execução de determinado pedaço de código

func funcao1() {
	fmt.Println("Executando função 1")
}

func funcao2() {
	fmt.Println("Executando função 2")
}

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada, o resultado será retornado")
	fmt.Println("Entrando na função para verificar se aluno está aprovado")

	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	return false
}

func main() {
	//executa primeiro a função 2 e dps 1

	// defer significa adiar, essa cláusula fala que seu código irá adiar a execução até o ultimo momento possível, como a função 1 não tem retorno, o último momento possível seria antes dela terminar. Se fosse uma função com retorno, o defer seria executado imediatamente antes do retorno ser dado.

	defer funcao1() // será executada depois da execução da função 2 e alunoAprovado
	funcao2()

	fmt.Println(alunoAprovado(5, 9))
}
