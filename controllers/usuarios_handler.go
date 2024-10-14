package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/emigoulart/digport-academy/model"
)

func CriarUsuarioHandler(w http.ResponseWriter, r *http.Request) {
	var usuario model.Usuario
	json.NewDecoder(r.Body).Decode(&usuario)

	err := usuario.Validar()
	if err != nil {
		fmt.Println("Usuário informado inválido:", err.Error())
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(model.Erro{Mensagem: err.Error()})
		return
	}

	err = model.CriaUsuario(usuario)
	if err != nil {
		fmt.Println("Erro ao criar usuário:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		fmt.Println("Usuário criado")
		w.WriteHeader(http.StatusCreated)
		return
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var usuario model.Usuario
	json.NewDecoder(r.Body).Decode(&usuario)

	username := usuario.Email
	senhatxt := usuario.Senha
	user, error := model.BuscaUsuarioPorEmail(username)
	if error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	hash := user.Senha

	error = model.ValidaLogin(hash, senhatxt)

	if error == nil {
		// Cria um token JWT com uma data de expiração
		GerarToken(w)

	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
	//
}

func BuscaUsuarioPorEmail(w http.ResponseWriter, r *http.Request) {

	email := r.URL.Query().Get("email")

	// http://localhost:8080/usuarios?email=emigoulart@gmail.com
	usuario, err := model.BuscaUsuarioPorEmail(email)
	if err != nil {
		fmt.Println("Erro ao buscar usuário:", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(usuario)
}
