package main

import "fmt"

func main() {
	var b [100]int
	for i := 0; i < 100; i++ {
		fmt.Printf("Digite o valor b[%d]: ", i)
		fmt.Scan(&b[i])
	}

	soma := 0
	for i := 0; i < 50; i++ {
		soma += (b[i] - b[99-i]) * (b[i] - b[99-i]) * (b[i] - b[99-i])
	}

	fmt.Println("Resultado do somatório:", soma)
}
