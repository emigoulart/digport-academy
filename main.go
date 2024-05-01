package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/emigoulart/digport-academy/model"

	_ "github.com/lib/pq"
)

func main() {

	fmt.Println("Bem vindo(a) à loja DigPort!")
	//fmt.Printf("Esse é o catálogo de produtos disponíveis:\n%+v\n", criaEstoque())
	// uuid := uuid.NewString()
	// fmt.Println(uuid)
	db := conectaBancoDados()

	resultado, err := db.Query("SELECT * FROM produtos")

	if err != nil {
		panic(err.Error())
	}

	p := model.Produto{}

	for resultado.Next() {

		var quantidade int
		var id, nome, descricao, imagem string
		var preco float64
		err = resultado.Scan(&id, &nome, &descricao, &imagem, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.ID = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Imagem = imagem
		p.QuantidadeEmEstoque = quantidade

		fmt.Println(p)
	}

	defer db.Close()
}

func conectaBancoDados() *sql.DB {
	connStr := "user=postgres dbname=digport_loja password=digport host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
