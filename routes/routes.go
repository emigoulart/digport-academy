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
	route.HandleFunc("/produto", controllers.AdicionaProdutoHandler).Methods("POST")
	route.HandleFunc("/produto", controllers.RemoveProdutoHandler).Methods("DELETE")
	http.ListenAndServe(":8085", route)
}
