package main

import (
	"fmt"

)

func main() {
	i := 1
	for i <= 10 { //Não existe while no GO
			fmt.Println(i)
			i ++  // para valor ir subindo até 10
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	// inicializa a variável i como 0
	// O loop continuará enquanto i for menor ou igual a 20
	// Após cada iteração é incrementada +2
	
	fmt.Println("\n-----------------------------------------")
	fmt.Println("\nMisturando para mostrar pares e Ímpares...")
	for i := 1; i <= 10; i++ {
		  if i % 2 == 0 {
				 fmt.Println(i, " - Par ")
			} else {
				fmt.Println(i, " - Impar ")
			}
	}
}
