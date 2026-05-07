package main

import "fmt"

func main() {
	var v [30]int
	var novo [30]int

	for i := 0; i < 30; i++ {
		fmt.Printf("Digite v[%d]: ", i)
		fmt.Scan(&v[i])
		if i%2 == 0 {
			novo[i] = v[i] * 2
		} else {
			novo[i] = v[i] * 3
		}
	}

	fmt.Println("Novo vetor:", novo)
}
