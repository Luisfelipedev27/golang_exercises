package main

import "fmt"
// GOLANG É FORTEMENTE TIPADA. A VARIAVEL SEMPRE TERÁ UM TIPO, ESPECIFICANDO OU NAO
// NO GO, VOCE PRECISA USAR A VARIAVEL QUE VOCE ESTÁ DECLARANDO
func main() {
	var variavel1 string = "Variavel 1"
	variavel2 := "Variavel 2" //Mesmo nao declarando o tipo, ela tem um tipo. Pelo valor da variável, ele advinha qual é o tipo
	//Nome técnico: Inferencia de tipo
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "exemplo3"
		variavel4 string = "exemplo4"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variavel 5", "Variavel 6"
	fmt.Println(variavel5, variavel6)

	//CONSTANTE
	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
	//Posso trocar/inverter o valor das variaveis
}
