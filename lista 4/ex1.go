package main

import (
	"fmt"
)

func main() {
	var numeros [10]int
	encontrou := false

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o %dº número: ", i+1)
		fmt.Scan(&numeros[i])
	}

	fmt.Println("\nNúmeros maiores que 50:")

	for i, valor := range numeros {
		if valor > 50 {
			fmt.Printf("Número: %d | Posição: %d\n", valor, i)
			encontrou = true
		}
	}

	if !encontrou {
		fmt.Println("Nenhum número maior que 50 foi encontrado.")
	}
}
