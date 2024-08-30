package main

import "github.com/emigoulart/digport-academy/db"

func main() {
	db.InitDB()
	StartServer()
}
