package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	connStr := "user=postgres dbname=digport_loja password=digport host=localhost sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	//defer DB.Close()
	criaTabelas()
}

// username TEXT NOT NULL UNIQUE,
func criaTabelas() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS USUARIO (
		id SERIAL PRIMARY KEY,
		nome VARCHAR,
		telefone VARCHAR,
		endereco VARCHAR,
		email VARCHAR NOT NULL,
		senha VARCHAR NOT NULL
	);`
	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Erro ao criar tabela usuario.")
	} else {
		fmt.Printf("tabela usuario ok")
	}

}
