package main

import "github.com/emigoulart/digport-academy/model"

func criaEstoque() []model.Produto {
	produtos := []model.Produto{
		{
			ID:                  "1",
			Nome:                "Revista Recreio",
			Preco:               19.90,
			Descricao:           "Revista Recreio",
			Imagem:              "revista.jpg",
			QuantidadeEmEstoque: 7,
		},

		{
			ID:                  "2",
			Nome:                "Revista Capricho",
			Preco:               14.0,
			Descricao:           "Revista Capricho",
			Imagem:              "revista.jpg",
			QuantidadeEmEstoque: 5,
		},

		{
			ID:                  "3",
			Nome:                "Revista Superinteressante",
			Preco:               18.9,
			Descricao:           "Revista Superinteressante",
			Imagem:              "revista.jpg",
			QuantidadeEmEstoque: 3,
		},

		{
			ID:                  "4",
			Nome:                "Devoradores de Estrelas",
			Preco:               75.0,
			Descricao:           "Devoradores de Estrelas",
			Imagem:              "livro.jpg",
			QuantidadeEmEstoque: 1,
		},

		{
			ID:                  "5",
			Nome:                "O Poder do Hábito",
			Preco:               66.0,
			Descricao:           "O Poder do Hábito",
			Imagem:              "livro.jpg",
			QuantidadeEmEstoque: 12,
		},

		{
			ID:                  "6",
			Nome:                "A menina que roubava livros",
			Preco:               39.90,
			Descricao:           "A menina que roubava livros",
			Imagem:              "livro.jpg",
			QuantidadeEmEstoque: 9,
		},

		{
			ID:                  "7",
			Nome:                "O Pequeno Príncipe",
			Preco:               19.90,
			Descricao:           "O Pequeno Príncipe",
			Imagem:              "livro.jpg",
			QuantidadeEmEstoque: 4,
		},
	}

	estoque := make([]model.Produto, len(produtos))

	return estoque
}
