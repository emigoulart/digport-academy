package main

import (
	"fmt"
)

func main() {
	fmt.Println("Bem vindo(a) à loja DigPort!")
	fmt.Printf("Esse é o catálogo de produtos disponíveis:\n%+v\n", criaEstoque())
}
