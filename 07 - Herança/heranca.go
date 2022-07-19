package main

import "fmt"

// todo estudante é uma pessoa

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	//nao precisamos declarar tudo do primeiro struct, estariamos sendo redundantes
	pessoa    // basta chamar o nome do struct, assim podemos acessar normalmente os campos do primeiro struct, pessoa já é um tipo, não precisa declarar um tipo.
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"João", "Pedro", 20, 189}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1.nome) // nao precisa colocar estudante.pessoa.nome, somente estudante.nome, pois ele já herda
	// o go não te permite criar variáveis e não usar, porém structs pode, pois são tipos
}
