package main

import (
	"fmt"
	"module/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {

	fmt.Println("Escrevendo da main")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("teste@gmail.com")
	fmt.Println(erro)
	// Ao referenciar, colocar somente o nome após a última barra, exemplo: module/auxiliar -> auxiliar
	// github.com/badoux/checkmail = checkmail
}

//go mod tidy = remove dependências não utilizadas

// go build = atualizar módulo
// go install = faz a mesma coisa do build, porém troca o diretório do módulos, colocando na raíz.

// No Go, como não é orientada a objetos, para saber o que é uma variável, função etc, se essas coisas são públicas, é por conta da primeira letra delas, se uma função começa minúscula, ela só vai ser visível no pacote que ela está. Se for maiúsculo, ela será pública.
