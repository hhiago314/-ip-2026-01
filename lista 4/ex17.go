package main

import "fmt"

func ehPrimo(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var v [10]int
	for i := 0; i < 10; i++ {
		fmt.Printf("Digite v[%d]: ", i)
		fmt.Scan(&v[i])
	}

	for i, val := range v {
		if ehPrimo(val) {
			fmt.Printf("Número primo %d na posição %d\n", val, i)
		}
	}
}
