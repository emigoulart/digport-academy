package main

import "fmt"

func main() {

	var nome string
	fmt.Printf("Olá, digite o ser nome: ")

	fmt.Scan(&nome)

	fmt.Printf("Bem vinda, %s", nome)

}
