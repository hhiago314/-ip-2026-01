package main

import "fmt"

func main() {
	h := 0.0
	sinal := 1.0
	for i := 1; i <= 50; i++ {
		h += sinal / float64(i*2-1)
		sinal *= -1
	}
	fmt.Printf("H = %.6f\n", h)
}
