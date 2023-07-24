package main

import (
	"fmt"
	"golang-rest-api/infra/database"
	"golang-rest-api/routes"
)

func main() {
	fmt.Println("Iniciando banco de dados")
	database.Connect()

	fmt.Println("Iniciando servidor")
	routes.HandleRequest()
}
