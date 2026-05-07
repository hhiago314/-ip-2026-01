package main

import (
	"fmt"
	"math"
)

func main() {
	var v [15]int
	var resultado [15]float64

	for i := 0; i < 15; i++ {
		fmt.Printf("Digite o %dº número: ", i+1)
		fmt.Scan(&v[i])
		if v[i] < 0 {
			resultado[i] = -1
		} else {
			resultado[i] = math.Sqrt(float64(v[i]))
		}
	}

	fmt.Println("Resultados:", resultado)
}
