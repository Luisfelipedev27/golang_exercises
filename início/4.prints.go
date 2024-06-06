package main

import "fmt"

func main() {
	fmt.Print("MESMA")
	fmt.Print(" LINHA.")

x := 3.14

// fmt.Println("O valor de X é", X)
xs := fmt.Sprint(x) //Sprint vem do pacote fmt... ele converte a variável x em string. Esse resultado é atribuído a variável xs
// Sprint converte qualquer tipo de dadoo em string. É mais versátil

fmt.Println("O valor de x é " + xs)
fmt.Println("O valor de x é", + x) // Aqui naão é uma string, então é necessário uma vírgula

fmt.Printf("O valor de x é %.2f.", x) // %.2f usado para formatar impressão de um número ponto flutuante. o .2 indica 2 casas decimais

a := 1
b := 1.9999
c := false
d := "opa"

fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
}

// %d é usado para formatar um número inteiro (a).
// %f é usado para formatar um número de ponto flutuante com seis casas decimais (b).
// %.1f é usado para formatar um número de ponto flutuante com uma casa decimal (b).
// %t é usado para formatar um valor booleano (c).
// %s é usado para formatar uma string (d).
