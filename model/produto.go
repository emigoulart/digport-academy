package model

import (
	"database/sql"
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
	defer db.Close()

	resultado, err := db.Query("SELECT * FROM produto")
	if err != nil {
		panic(err.Error())
	}

	produtos := []Produto{}

	for resultado.Next() {
		err = resultado.Scan(&id, &nome, &preco, &descricao, &imagem, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		var produto = populaProduto()

		produtos = append(produtos, produto)
	}
	return produtos
}

func BuscaProdutoPorNome(nomeProduto string) Produto {
	db := db.ConectaBancoDados()
	defer db.Close()

	res := db.QueryRow("SELECT * FROM produtos where nome = $1", nomeProduto)

	err := res.Scan(&id, &nome, &preco, &descricao, &imagem, &quantidade)
	if err == sql.ErrNoRows {
		fmt.Printf("Produto não encontrado %s\n", nomeProduto)

	} else if err != nil {
		panic(err.Error())
	}

	var produto1 = populaProduto()

	return produto1
}

func populaProduto() Produto {
	var produto1 Produto
	produto1.ID = id
	produto1.Nome = nome
	produto1.Descricao = descricao
	produto1.Preco = preco
	produto1.Imagem = imagem
	produto1.QuantidadeEmEstoque = quantidade
	return produto1

}

func CriaProduto(prod Produto) error {
	//nome, descricao string, preco float64, image string, quantidade int
	if produtoCadastrado(prod.Nome) {
		fmt.Printf("Produto já cadastrado: %s\n", prod.Nome)
		return fmt.Errorf("Produto já cadastrado")
	}
	id := uuid.NewString()
	nome := prod.Nome
	preco := prod.Preco
	descricao := prod.Descricao
	imagem := prod.Imagem
	quantidade := prod.QuantidadeEmEstoque

	strInsert := "INSERT INTO produtos VALUES($1, $2, $3, $4, $5, $6)"

	db := db.ConectaBancoDados()
	defer db.Close()

	result, err := db.Exec(strInsert, id, nome, strconv.FormatFloat(preco, 'f', 1, 64), descricao, imagem, strconv.Itoa(quantidade))

	if err != nil {
		panic(err.Error())
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Produto %s criado com sucesso (%d row affected)\n", id, rowsAffected)

	return nil
}

func produtoCadastrado(nomeProduto string) bool {

	prod := BuscaProdutoPorNome(nomeProduto)

	return prod.Nome == nomeProduto
}

func RemoveProduto(id string) error {
	db := db.ConectaBancoDados()
	defer db.Close()

	result, err := db.Exec("DELETE FROM produtos WHERE id = $1", id)
	if err != nil {
		panic(err.Error())
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("id: ", id)
	if rowsAffected == 0 {
		return fmt.Errorf("produto não encontrado")
	}

	fmt.Printf("Produto %s deletado com sucesso (%d row affected)\n", id, rowsAffected)

	return nil
}

func UpdateProduto(prod Produto) error {
	db := db.ConectaBancoDados()
	defer db.Close()

	id := prod.ID
	nome := prod.Nome
	//preco := prod.Preco
	descricao := prod.Descricao
	//imagem := prod.Imagem
	//quantidade := prod.QuantidadeEmEstoque

	result, err := db.Exec("UPDATE produtos SET nome= $1, descricao= $2 where id= $3", nome, descricao, id)
	if err != nil {
		panic(err.Error())
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	if rowsAffected == 0 {
		return fmt.Errorf("produto não encontrado")
	}

	fmt.Printf("Produto %s atualizado com sucesso (%d row affected)\n", id, rowsAffected)

	return nil
}
