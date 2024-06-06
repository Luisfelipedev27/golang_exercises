package main

import (
	"modulo/auxiliar" // Primeiro chamando o módulo e depois o pacote
	"fmt"
	"github.com/badoux/checkmail" // Importando um pacote de terceiros
)

func main(){
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(erro) // Existe erro ?
}

// usei o go build para compilar o arquivo main.go
// ./modulo para executar o arquivo compilado
// Se eu nao quiser utilizar mais o pacote, removo do modulo com ->  go mod tidy
