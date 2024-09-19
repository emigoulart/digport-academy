package routes

import (
	"net/http"

	"github.com/emigoulart/digport-academy/controllers"
	"github.com/gorilla/mux"
)

func HandleRequests() {
	route := mux.NewRouter()
	route.HandleFunc("/produtos", controllers.BuscaProdutosHandler).Methods("GET")
	route.HandleFunc("/produto", controllers.BuscaProdutoPorNomeHandler).Methods("GET")
	route.HandleFunc("/produto", controllers.CriaProdutoHandler).Methods("POST")
	route.HandleFunc("/produto/{id}", controllers.RemoveProdutoHandler).Methods("DELETE")
	route.HandleFunc("/produto", controllers.AtualizaProdutoHandler).Methods("PUT")

//Usuario
	route.HandleFunc("/usuarios", controllers.CriarUsuarioHandler).Methods("POST")
	//route.HandleFunc("/usuarios",controllers.BuscaUsuarioPorEmail).Methods("GET")
	//route.HandleFunc("/usuarios/login",controllers.LoginHandler).Methods("POST")

	http.ListenAndServe(":8085", route)
}
