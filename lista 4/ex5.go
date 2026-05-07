package main

import "fmt"

func main() {
	var v [10]int
	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o %dº número: ", i+1)
		fmt.Scan(&v[i])
	}

	menor := v[0]
	pos := 0
	for i, val := range v {
		if val < menor {
			menor = val
			pos = i
		}
	}
	fmt.Printf("O menor elemento do vetor é %d e sua posição é %d\n", menor, pos)
}
