package model

import (
	"github.com/emigoulart/digport-academy/db"

	"golang.org/x/crypto/bcrypt"
)

type Usuario struct {
	ID       int64  `json:"id"`
	Nome     string `json:"nome"`
	Senha    string `json:"senha"`
	Email    string `json:"email"`
	Telefone string `json:"telefone"`
	Endereco string `json:"endereco"`
}

func CriaUsuario(usuario Usuario) error {
	password, err := hashPassword(usuario.Senha)
	if err != nil {
		return err
	}

	db := db.ConectaBancoDados()
	defer db.Close()

	_, err = db.Exec("INSERT INTO usuario (nome, senha, email, telefone, endereco) VALUES ($1, $2, $3, $4, $5)",
		usuario.Nome, password, usuario.Email, usuario.Telefone, usuario.Endereco)
	if err != nil {
		return err
	}
	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
