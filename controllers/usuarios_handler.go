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

	err := model.CriaUsuario(usuario)
	if err != nil {
		fmt.Println("Erro ao criar usuário:", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		fmt.Println("Usuário criado")
		w.WriteHeader(http.StatusCreated)
	}
}
