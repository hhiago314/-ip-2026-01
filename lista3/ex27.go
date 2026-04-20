package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Print("Digite x em radianos: ")
	fmt.Scan(&x)

	cosSerie := 1.0
	sinal := -1.0
	pot := x * x
	fat := 2.0

	for i := 2; i <= 20; i += 2 {
		cosSerie += sinal * math.Pow(x, float64(i)) / float64(fatorial(i))
		sinal *= -1
	}

	fmt.Printf("Cosseno pela série: %.6f\n", cosSerie)
	fmt.Printf("Cosseno pela função: %.6f\n", math.Cos(x))
	fmt.Printf("Diferença: %.6f\n", math.Cos(x)-cosSerie)
}

func fatorial(n int) int {
	fat := 1
	for i := 2; i <= n; i++ {
		fat *= i
	}
	return fat
}
