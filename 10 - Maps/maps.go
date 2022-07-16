package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario["nome"]) //nao chamamos por usuario.nome, e sim usuario["nome"]

	usuario2 := map[string]map[string]string{ //aninhar maps
		"nome": {
			"primeiro": "Joao",
			"ultimo":   "pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)
}
