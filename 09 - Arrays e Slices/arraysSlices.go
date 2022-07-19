package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int
	array1[2] = 3
	fmt.Println(array1)

	array2 := [5]int{1, 2, 3, 4, 5}

	fmt.Println(array2)

	// não declarar tamanho no array ele vira um slice

	array3 := [...]int{1, 2, 3, 4, 5} //fixam tamanho do array de acordo com a quantidade de itens que vc passou dentro, não é dinamico apenas deixa com base no que vc colocou, não é muito utilizado em Go.
	fmt.Println(array3)

	//SLICE

	// Slice = fatia

	slice := []int{234, 234, 234, 2342}
	//muito usado em go, pois nn é limitado a quantidade de item
	fmt.Println(slice)

	//reflect = pacote que retorna o tipo da variável
	fmt.Println(reflect.TypeOf(slice))

	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 34) // pega o slice com todas informações que tinham e adiciona um item e retorna um slice novo

	fmt.Println(slice)

	slice2 := array2[1:3] // criar uma fatia (slice) a partir de um array que já existia, como se fosse um ponteiro
	fmt.Println(slice2)

	array2[2] = 30
	fmt.Println(slice2)

	// ARRAYS INTERNOS
	fmt.Println("-----------------------")
	slice3 := make([]int, 4, 5) // make = funcao para alocar espaço na memoria pra alguma coisa, como o slice nesse caso
	fmt.Println(slice3)

	slice3 = append(slice3, 3)
	slice3 = append(slice3, 7)
	slice3 = append(slice3, 0)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacidade, o Go duplica a capacidade caso estoure o tamanho, ele cria um novo array para você referenciar, e dobra o tamanho dele

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4)) // quando vc não passa o ultimo parametro no make, ele vai definir que a capacidade do slice é igual ao tamanho dele, porém se estourar ainda vai duplicar

	//array = lista com tamanho fixo
	//slice = lista sem tamanho fixo
}
