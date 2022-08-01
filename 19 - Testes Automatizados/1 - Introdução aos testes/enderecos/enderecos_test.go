//TESTE DE UNIDADE = TESTA A MENOR UNIDADE DO SEU CÓDIGO, NESSE CASO UMA FUNÇÃO

package enderecos

import (
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

//sintaxe de uma função de teste, obrigatoriamente começa com Test, uma boa prática é colocar exatamente o nome da função, e depois recebe um parametro que é um ponteiro do tipo T que está dentro do pacote testing.

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel() //rodar em paralelo

	cenarioDeTeste := []cenarioDeTeste{

		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Estrada Qualquer", "Estrada"},
		{"RUA ABC", "Rua"},
		{"AVENIDA PAULISTA", "Avenida"},
		{"Praça das Rosas", "Tipo Inválido"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenarioDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				retornoRecebido,
				cenario.retornoEsperado,
			)
		}
	}

}

func TestQualquer(t *testing.T) {
	t.Parallel() //roda em paralelo somente testes que possuem t.Parallel()
	if 1 > 3 {
		t.Errorf("O teste falhou")
	}
}
