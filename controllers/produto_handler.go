package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/emigoulart/digport-academy/model"
	"github.com/gorilla/mux"
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

func BuscaProdutoPorIdHandler(w http.ResponseWriter, r *http.Request) {
	// http://localhost:8080/produto/{id}
	fmt.Print("BuscaProdutoPorIdHandler")
	vars := mux.Vars(r)
	id := vars["id"]
	produto, err := model.BuscaProdutoPorId(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(produto)

}

func CriaProdutoHandler(w http.ResponseWriter, r *http.Request) {
	var produto model.Produto
	json.NewDecoder(r.Body).Decode(&produto)

	err := model.CriaProduto(produto)
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func RemoveProdutoHandler(w http.ResponseWriter, r *http.Request) {
	// implementation of the RemoveProdutoHandler function
	// the function should receive a request and remove a product from the database
	// the product to be removed should be passed as a parameter in the request body
	// the function should return a status code 204 if the product was removed successfully, no content
	// or a status code 404 if the product was not found
	id := mux.Vars(r)["id"]
	err := model.RemoveProduto(id)
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func AtualizaProdutoHandler(w http.ResponseWriter, r *http.Request) {
	// implementation of AtualizaProdutoHandler
	var produto model.Produto
	json.NewDecoder(r.Body).Decode(&produto)
	err := model.UpdateProduto(produto)
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Println(produto.ID)
	w.WriteHeader(http.StatusOK)

}
