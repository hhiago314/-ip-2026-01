package main

import (
	"fmt"
)

func power(x, n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= x
	}
	return result
}

func main() {
	var x, n int
	fmt.Print("Digite o valor de x: ")
	fmt.Scan(&x)
	fmt.Print("Digite o valor de n: ")
	fmt.Scan(&n)

	fmt.Printf("%d^%d = %d\n", x, n, power(x, n))
}
