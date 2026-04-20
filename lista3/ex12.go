package main

import "fmt"

func main() {
	var x float64
	fmt.Print("Digite um número real X: ")
	fmt.Scan(&x)

	s := x
	fat := 1.0
	pot := x

	for i := 1; i <= 20; i++ {
		fat *= float64(i)
		pot *= x
		s += pot / fat
	}

	fmt.Printf("Série = %.6f\n", s)
}
