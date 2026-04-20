package main

import "fmt"

func main() {
	s := 0.0
	num := 1
	den := 225
	for i := 0; i < 15; i++ {
		s += float64(num) / float64(den)
		num *= 2
		den -= 29
	}
	fmt.Printf("S = %.6f\n", s)
}
