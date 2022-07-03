package auxiliar

import "fmt"

func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar")

	//Por estar no mesmo pacote, você consegue chamar uma função de outro arquivo, e sem a necessidade de reescrever o "auxiliar.Escrever()" como acontece na main.go

	escrever2()
}

// Por padrão no Go, quando exportar uma função você deve ter um comentário em cima dela, falando o que ela faz. Por uma questão de boa prática. (O Go gera um warning, e não um erro)
