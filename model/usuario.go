package model

import (
	"database/sql"
	"fmt"

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

func (u Usuario) Validar() error {
	if u.Nome == "" {
		return fmt.Errorf("nome não pode ser vazio")
	}
	if u.Email == "" {
		return fmt.Errorf("email não pode ser vazio")
	}
	if u.Senha == "" {
		return fmt.Errorf("senha não pode ser vazia")
	}
	return nil
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

func ValidaLogin(hash string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return fmt.Errorf("invalid password")
	}
	return nil
}

func BuscaUsuarioPorEmail(email string) (*Usuario, error) {
	db := db.ConectaBancoDados()
	defer db.Close()

	var usuario Usuario
	err := db.QueryRow("SELECT id, nome, senha, email, telefone, endereco FROM usuario WHERE email = $1", email).Scan(
		&usuario.ID, &usuario.Nome, &usuario.Senha, &usuario.Email, &usuario.Telefone, &usuario.Endereco,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Usuario não encontrado %s", email)
		}
		return nil, err
	}

	return &usuario, nil
}
