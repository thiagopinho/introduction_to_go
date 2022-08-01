package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"` // que chaves eles vão se tornar
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	//json.Marshal() converter struct to json ou map to json, unMarshal é o contrario
	c := cachorro{"Rex", "Dálmata", 3}
	fmt.Println(c)

	cachorroEmJSON, erro := json.Marshal(c)

	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorroEmJSON)

	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

	c2 := map[string]string{
		"nome": "Tob",
		"raca": "Poodle",
	}

	cachorro2EmJSON, erro := json.Marshal(c2)

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorro2EmJSON)

	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))
}
