package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/emigoulart/digport-academy/model"
)

func BuscaProdutosHandler(w http.ResponseWriter, r *http.Request) {
	produtos := model.BuscaTodosProdutos()
	json.NewEncoder(w).Encode(produtos)

}

func BuscaProdutoPorNomeHandler(w http.ResponseWriter, r *http.Request) {
	// o parametro será passado como parte da url
	// http://localhost:8080/produto?nome=Revista Capricho

	//obtem o parametro nome da url
	nome := r.URL.Query().Get("nome")
	produto := model.BuscaProdutoPorNome(nome)
	json.NewEncoder(w).Encode(produto)

}
