package main

import "fmt"

func main() {
	var id, peso int
	var idMax, idMin, pesoMax, pesoMin int

	for i := 1; i <= 90; i++ {
		fmt.Printf("Digite id e peso do boi %d: ", i)
		fmt.Scan(&id, &peso)

		if i == 1 || peso > pesoMax {
			pesoMax = peso
			idMax = id
		}
		if i == 1 || peso < pesoMin {
			pesoMin = peso
			idMin = id
		}
	}

	fmt.Printf("Mais gordo: id=%d peso=%d\n", idMax, pesoMax)
	fmt.Printf("Mais magro: id=%d peso=%d\n", idMin, pesoMin)
}
