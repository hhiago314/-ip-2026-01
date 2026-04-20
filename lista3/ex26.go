package main

import "fmt"

func main() {
	soma := 0.0
	sinal := 1.0
	for i := 0; i < 20; i++ {
		num := 100 - i
		fat := 1
		for j := 1; j <= i; j++ {
			fat *= j
		}
		soma += sinal * float64(num) / float64(fat)
		sinal *= -1
	}
	fmt.Printf("Soma = %.6f\n", soma)
}
