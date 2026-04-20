package main

import (
	"fmt"
)

func senoMacLaurin(x float64) float64 {
	resultado := 0.0
	termo := x
	sinal := 1.0
	fat := 1.0
	pot := x

	for i := 1; i <= 10; i++ {
		resultado += sinal * pot / fat
		sinal *= -1
		pot *= x * x
		fat *= float64((2 * i) * (2*i + 1))
	}
	return resultado
}

func main() {
	for a := 0.0; a <= 6.3; a += 0.1 {
		fmt.Printf("A=%.1f -> Sen(A)=%.6f\n", a, senoMacLaurin(a))
	}
}
