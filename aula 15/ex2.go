package main

import (
	"fmt"
)

func sumArray(arr []float64) float64 {
	sum := 0.0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func main() {
	var n int
	fmt.Print("Digite o tamanho do array: ")
	fmt.Scan(&n)

	arr := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Digite o valor %d: ", i+1)
		fmt.Scan(&arr[i])
	}

	fmt.Println("Soma =", sumArray(arr))
}
