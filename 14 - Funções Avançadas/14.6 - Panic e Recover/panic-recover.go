package main

import "fmt"

func recuperarExecucao() {
	//fmt.Println("tentativa de recuperar execucao")
	if r := recover(); r != nil {
		fmt.Println("recuperada com sucesso")
	}
}

func alunoAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	// Interrompe o fluxo normal do programa e para tudo, quando chamamos, ela entra em "panico", ela vai chamar todas as funções com defer, seu programa morre, para de executar.

	panic("A MEDIA É EXATAMENTE 6")
}

func main() {
	fmt.Println(alunoAprovado(6, 6)) // entra em panico e mata a execução do programa se a média é 6
	fmt.Println("Pós execução")
}
