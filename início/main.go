package main

import "fmt"

func main() {
	resultado := somar(3, 4)
	imprimir(resultado)
}

func somar(a int, b int) int {  // esse int significa que sempre retorna um inteiro. Tipando o retorno
	return a + b
}

func imprimir(valor int) {
	fmt.Println(valor)
}
