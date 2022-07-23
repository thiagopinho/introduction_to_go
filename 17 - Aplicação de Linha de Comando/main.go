package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {

	fmt.Println("Ponto de partida")
	aplicacao := app.Gerar()

	erro := aplicacao.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}

	//parametro padrão para que os comandos do sistema operacional sejam reconhecidos pela linha de comando

}
