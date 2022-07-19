package main

import "fmt"

func main() {
	func() {
		fmt.Println("Olá mundo!")
	}() // falando pro GO que isso é uma função anonima, assim que terminar de declarar, execute.

	//com parametro
	func(texto string) {
		fmt.Println(texto)
	}("Passando texto") // Como a função anonima não tem nome e não podemos chama-la externamente, é assim que passamos no parametro

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido => %s", texto) //SprintF = usado para concatenar informações, concatenar com outros tipos de dados %s = string %d = numero
	}("Passando texto")

	fmt.Println(retorno)
}
