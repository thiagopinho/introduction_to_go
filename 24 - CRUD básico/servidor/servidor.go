package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")

	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}
	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o ID inserido"))
		return
	}

	//STATUS CODE

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso!: ID: %d", idInserido)))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
	}
	defer db.Close()

	linhas, erro := db.Query("select * from usuarios")

	if erro != nil {
		w.Write([]byte("Erro ao buscar os usuários"))
	}
	defer linhas.Close()

	var usuarios []usuario
	//para cada linha fazer uma iteração
	for linhas.Next() {
		var usuario usuario //populando o slice de usuarios

		if erro != linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email) {
			w.Write([]byte("Erro ao escanear o usuário"))
			return //sempre apos um w.Write precisamos dar um retorno pra ele nao continuar executando a funcao
		}

		usuarios = append(usuarios, usuario)

	}
	w.WriteHeader(http.StatusOK)

	//resposta via http

	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter os usuários para JSON"))
	}

}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

}
