package main

import (
	"fmt"
	"time" // Fornece funcionalidades para medir e exibir o tempo
)

func main() {
	fmt.Println("Strings:", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("-------------------------- Novo Bloco----------------------")
	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas:", d1 == d2 )           // Ambas Fazem a mesma coisa, comparando valores
	fmt.Println("Datas:", d1.Equal(d2))

	type Pessoa struct {    // Define uma estrutura chamada Pessoa, cria duas instâncias dessa estrutura e, em seguida, compara as duas instâncias.
		Nome string
		Idade int
		Endereço string
	}
												//  Struct é uma maneira de agrupar/encapsular várias variáveis em um único tipo de dados.
	p1 := Pessoa{Nome: "João"}
	p2 := Pessoa{Nome: "João"}
	y := Pessoa{"João", 30, "Rua ABC"}

	fmt.Println("Pessoas:", p1 == p2)
	fmt.Print(y)
}

