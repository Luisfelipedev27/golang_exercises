package main

import "fmt"

func imprimirResultado(nota float64) {
	if nota >= 7 {
			fmt.Println("Aprovado com nota", nota)
	} else {
			fmt.Println("Reprovado com nota", nota)
	}
}

func notaParaConceito(n float64) string {
		if n >= 9 && n <= 10 {
				return "A"
		} else if n >= 8 && n < 9 {
				return "B"
		} else if n >= 5 && n < 8 {
				return "C"
		} else {
			return "E"
		}
}

func main() {
	imprimirResultado(7.3)
	imprimirResultado(5.1)

	fmt.Println("---------------------------")
	
	fmt.Println((notaParaConceito(9.8)))
	fmt.Println((notaParaConceito(6.9)))
	fmt.Println((notaParaConceito(4.4)))
}
