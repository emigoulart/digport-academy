package model

import (
	"fmt"
	"strconv"

	"github.com/emigoulart/digport-academy/db"
	"github.com/google/uuid"
)

type Produto struct {
	ID                  string  `json:"id"`
	Nome                string  `json:"nome"`
	Preco               float64 `json:"preco"`
	Descricao           string  `json:"descricao"`
	Imagem              string  `json:"imagem"`
	QuantidadeEmEstoque int     `json:"quantidadeEmEstoque"`
}

var id, nome string
var preco float64
var descricao, imagem string
var quantidade int

func BuscaTodosProdutos() []Produto {
	db := db.ConectaBancoDados()

	resultado, err := db.Query("SELECT * FROM produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for resultado.Next() {

		err = resultado.Scan(&id, &nome, &preco, &descricao, &imagem, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.ID = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Imagem = imagem
		p.QuantidadeEmEstoque = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos

}

func BuscaProdutoPorNome(nomeProduto string) Produto {
	db := db.ConectaBancoDados()

	res := db.QueryRow("SELECT * FROM produtos where nome = $1", nomeProduto)

	err := res.Scan(&id, &nome, &preco, &descricao, &imagem, &quantidade)

	if err != nil {
		panic(err.Error())
	}
	var p Produto
	p.ID = id
	p.Nome = nome
	p.Descricao = descricao
	p.Preco = preco
	p.Imagem = imagem
	p.QuantidadeEmEstoque = quantidade

	defer db.Close()
	return p

}

func CriaProduto(prod Produto) {
	//nome, descricao string, preco float64, image string, quantidade int
	db := db.ConectaBancoDados()
	id := uuid.NewString()
	nome := prod.Nome
	preco := prod.Preco
	descricao := prod.Descricao
	imagem := prod.Imagem
	quantidade := prod.QuantidadeEmEstoque

	result, err := db.Exec("INSERT INTO produtos VALUES($1, $2, $3, $4, $5, $6)", id, nome, strconv.FormatFloat(preco, 'f', 1, 64), descricao, imagem, strconv.Itoa(quantidade))
	if err != nil {
		panic(err.Error())
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		//http.Error(w, http.StatusText(500), 500)
		panic(err.Error())
		//return
	}

	fmt.Printf("Products %s created successfully (%d row affected)\n", id, rowsAffected)

	defer db.Close()
}
