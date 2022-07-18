package main

import "fmt"

func soma(numeros ...int) int {
	//soma entre N numeros, não queremos especificar quantos numeros serão, queremos ter essa flexibilidade
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total

}
func main() {
	totalDaSoma := soma(1, 2, 3, 4, 5)
	fmt.Println(totalDaSoma)

	escrever("Olá Mundo", 10, 2, 4, 6)
}

func escrever(texto string, numeros ...int) { // o variatico necessariamente precisa ser o ultimo
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}

}
