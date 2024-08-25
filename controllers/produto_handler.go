package controllers

import (
	"encoding/json"
	"fmt"
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
	nome := r.URL.Query().Get("nome")
	produto := model.BuscaProdutoPorNome(nome)
	json.NewEncoder(w).Encode(produto)

}

func CriaProdutoHandler(w http.ResponseWriter, r *http.Request) {
	var produto model.Produto
	json.NewDecoder(r.Body).Decode(&produto)

	error := model.CriaProduto(produto)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}

}

func RemoveProdutoHandler(w http.ResponseWriter, r *http.Request) {
	// implementation of the RemoveProdutoHandler function
	// the function should receive a request and remove a product from the database
	// the product to be removed should be passed as a parameter in the request body
	// the function should return a status code 204 if the product was removed successfully, no content
	// or a status code 404 if the product was not found
	var produto model.Produto
	json.NewDecoder(r.Body).Decode(&produto)
	error := model.RemoveProduto(produto.ID)
	if error != nil {
		fmt.Print(error)
		w.WriteHeader(http.StatusNotFound)
	} else {
		fmt.Println(produto.ID)
		w.WriteHeader(http.StatusNoContent)
	}
}

func AtualizaProdutoHandler(w http.ResponseWriter, r *http.Request) {
	// implementation of AtualizaProdutoHandler
	var produto model.Produto
	json.NewDecoder(r.Body).Decode(&produto)
	error := model.UpdateProduto(produto)
	if error != nil {
		fmt.Print(error)
		w.WriteHeader(http.StatusNotFound)
	} else {
		fmt.Println(produto.ID)
		w.WriteHeader(http.StatusOK)
	}

}
