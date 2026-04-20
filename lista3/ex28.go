package main

import "fmt"

func main() {
	s := 0.0
	sinal := 1.0
	for i := 1; i <= 51; i++ {
		s += sinal / float64(2*i-1)
		sinal *= -1
	}
	pi := s * 4
	fmt.Printf("Aproximação de π = %.6f\n", pi)
}
