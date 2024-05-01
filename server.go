package main

import (
	"github.com/emigoulart/digport-academy/routes"
)

func StartServer() {
	//http.HandleFunc("/produtos", controllers.BuscaProdutos)
	routes.HandleRequests()
}
