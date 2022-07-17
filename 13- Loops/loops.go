// No Go só existe For, não tem for in, while, do while etc.

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops")

	//enquanto for verdadeira repita, como o while

	// i := 0

	// for i < 10 {

	// 	i++
	// 	fmt.Println("Incrementando i")
	// 	time.Sleep(time.Second) //fazer o código "dormir, por um segundo"
	// }

	// fmt.Println(i)

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando J", j)
	// 	time.Sleep(time.Second)

	// }

	//clausula range = iterar em um array, slice, string, map, range é usado para percorrer nesses cenarios

	nomes := [3]string{"João", "Davi", "Nunca"}

	for indice, nome := range nomes { //ele da um range em nomes e cada indice e nome, fazemos algo, no caso do nome indice e nome, podemos usar qualquer nomenclatura, se não quiser por exemplo que o indice apareça, pode colocar um "_" underline para que vc possa visualizar apenas o nome, por padrão o index sempre vai ser o primeiro, logo precisa do underline
		fmt.Println(indice, nome)

	}
	//iterar string
	for indice, letra := range "palavra" {
		fmt.Println(indice, string(letra))
	}

	//iterar maps
	usuario := map[string]string{
		"Nome":  "leo",
		"sobre": "soo",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// não podemos fazer range em struct

}
