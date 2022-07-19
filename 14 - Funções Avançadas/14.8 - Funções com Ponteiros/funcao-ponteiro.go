package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}
func main() {
	// NÃO ALTERA O VALOR DA VAŔIÁVEL
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero) // o numero continua sendo 20 pois quando chamamos a função nós passamos uma cópia e jogamos na variável, a variável de cima (numero) fica-se inalterável

	// ALTERA O VALOR DA VAŔIÁVEL
	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1 //estamos jogando um valor novo dentro do endereço de memória e referenciamos o valor que já estava lá * -1 s
}
