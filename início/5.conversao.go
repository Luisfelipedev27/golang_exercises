//Conversão entre Tipos Básicos

package main

import (
	"fmt"
	"strconv" // é um pacote em Go que implementa conversões para e de representações de string de valores básicos de Go.
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota) //Transforma nota final em inteiro
	fmt.Println(notaFinal)

	//cuidado
	fmt.Println("Teste " + string(97)) //  No entanto, em Go, quando você converte um número inteiro em uma string, ele é tratado como um valor Unicode

	//correto -> // Itoa é uma abreviação de integer to ASCII. Isso converte numero inteiro em string
	fmt.Println("Teste " + strconv.Itoa(123))

  // string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)
	b, _ := strconv.ParseBool("true")
	if b {
      fmt.Println("Verdadeiro")
	}
}
