package main

import "fmt"

func main() {
	var A [10]int
	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o %dº número: ", i+1)
		fmt.Scan(&A[i])
	}

	contagem := make(map[int]int)
	for _, val := range A {
		contagem[val]++
	}

	for num, qtd := range contagem {
		if qtd > 1 {
			fmt.Printf("Número %d repetido %d vezes\n", num, qtd)
		}
	}
}
