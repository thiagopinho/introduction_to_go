package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

//criar como se fosse uma função. Mas ela está grudada a uma estrutura, todos usuarios tem um método salvar

//u = variavel que usamos para referenciar outros campos do mesmo usuário
func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuário 1", 5}
	usuario1.salvar()

	usuario2 := usuario{"Usuário 2", 20}
	usuario2.salvar()

	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
