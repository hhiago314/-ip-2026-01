package main

import (
	"fmt"
)

func main() {
	var numeros [10]int
	var pares []int
	var impares []int
	somaPares := 0

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o %dº número: ", i+1)
		fmt.Scan(&numeros[i])
	}

	for _, val := range numeros {
		if val%2 == 0 {
			pares = append(pares, val)
			somaPares += val
		} else {
			impares = append(impares, val)
		}
	}

	fmt.Println("Números pares:", pares)
	fmt.Println("Soma dos pares:", somaPares)
	fmt.Println("Números ímpares:", impares)
	fmt.Println("Quantidade de ímpares:", len(impares))
}
