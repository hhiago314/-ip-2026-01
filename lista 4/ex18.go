package main

import "fmt"

func main() {

	var vetor [10]int

	fmt.Println("Digite 10 números em ordem crescente:")

	for i := 0; i < 10; i++ {

		for {

			fmt.Printf("Digite o %dº número: ", i+1)
			fmt.Scan(&vetor[i])

			if i == 0 {
				break
			}

			if vetor[i] > vetor[i-1] {
				break
			}

			fmt.Println("O número deve ser maior que o anterior.")
		}
	}

	fmt.Println("\nVetor em ordem crescente:")

	for _, valor := range vetor {
		fmt.Print(valor, " ")
	}
}
