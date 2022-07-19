package main

import "fmt"

var n int

// A função init será executada antes da função main, somente uma por aquivo.
func init() {
	n = 10
	fmt.Println("Executando a função init")
}

func main() {

	fmt.Println("Função main sendo executada", n)
}
