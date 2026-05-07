package main

import "fmt"

func main() {
	var v1, v2 [10]int
	var result []int

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite v1[%d]: ", i)
		fmt.Scan(&v1[i])
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("Digite v2[%d]: ", i)
		fmt.Scan(&v2[i])
	}

	for i := 0; i < 10; i++ {
		result = append(result, v1[i], v2[i])
	}

	fmt.Println("Vetor intercalado:", result)
}
