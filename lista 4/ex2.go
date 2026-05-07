package main

import (
	"fmt"
)

func main() {
	var v1 [10]int
	var v2 [5]int

	// Lendo primeiro vetor
	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o %dº número do primeiro vetor: ", i+1)
		fmt.Scan(&v1[i])
	}

	// Lendo segundo vetor
	for i := 0; i < 5; i++ {
		fmt.Printf("Digite o %dº número do segundo vetor: ", i+1)
		fmt.Scan(&v2[i])
	}

	// Calculando soma de todos elementos do segundo vetor
	somaV2 := 0
	for _, val := range v2 {
		somaV2 += val
	}

	// Vetores resultantes
	var pares []int
	var impares []int

	for _, val := range v1 {
		if val%2 == 0 {
			pares = append(pares, val+somaV2)
		} else {
			impares = append(impares, val+somaV2)
		}
	}

	fmt.Println("Vetor resultante dos pares:", pares)
	fmt.Println("Vetor resultante dos ímpares:", impares)
}
