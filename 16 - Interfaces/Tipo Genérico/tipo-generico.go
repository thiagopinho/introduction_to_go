package main

import "fmt"

//pode receber um parametro que atenda essa interface, como não tem nada, tudo atende, uma string, um int, booleano etc.
func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {

	//recebe todos os tipos

	generica("String")
	generica(2)
	generica(false)

	// Exemplo do que não fazer: Pois fica muito bagunçado

	mapa := map[interface{}]interface{}{
		1:            "String",
		float32(100): true,
		"String":     "String",
	}

	fmt.Println(mapa)
}

//Há funções que fazem sentido receber um genérico para interagir com qualquer coisa. Exemplo de uso: A função do PrintLn recebe vários tipos de dados, pois uma função que printa na tela precisa ter uma flexibilidade.
