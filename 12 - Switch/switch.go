package main

import "fmt"

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(1)
	fmt.Println(dia)
}

func diaDaSemana(numero int) string { //recebe numero, retorna string

	var diaDaSemana string

	switch numero {
	case 1:
		diaDaSemana = "Domingo"
		fallthrough //joga seu código pra dentro da próxima condição, ele não avalia, joga direto, essa função não é mt utilizada
	case 2:
		diaDaSemana = "Segunda"
	case 3:
		diaDaSemana = "Terça"
	case 4:
		diaDaSemana = "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
	return diaDaSemana
}
