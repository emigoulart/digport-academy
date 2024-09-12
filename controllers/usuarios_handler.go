package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/emigoulart/digport-academy/model"
	"github.com/golang-jwt/jwt"
)

// go get github.com/golang-jwt/jwt
// https://jwt.io/#debugger-io

var jwtKey = []byte("secret") // secret para a validação da assinatura

func CriarUsuarioHandler(w http.ResponseWriter, r *http.Request) {
	var usuario model.Usuario
	json.NewDecoder(r.Body).Decode(&usuario)

	err := usuario.Validar()
	if err != nil {
		fmt.Println("Usuário informado inválido:", err)
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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Verifica as credenciais do usuário
	var usuario model.Usuario
	json.NewDecoder(r.Body).Decode(&usuario)

	username := usuario.Email
	password := usuario.Senha

	// http://localhost:8080/usuarios?email=emigoulart@gmail.com
	userDB, err := model.BuscaUsuarioPorEmail(username)

	if err != nil {
		fmt.Println("Erro ao buscar usuário:", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Verifica se a senha é valida
	err = model.ValidaLogin(userDB.Senha, password)

	if err == nil {
		// Cria um token JWT com uma data de expiração
		expirationTime := time.Now().Add(5 * time.Minute)
		claims := &jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Retorna o token JWT para o cliente
		w.Write([]byte(tokenString))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}
