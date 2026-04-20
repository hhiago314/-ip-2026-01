package main

import "fmt"

func main() {
	var N int
	fmt.Print("Digite N: ")
	fmt.Scan(&N)

	soma := 0
	valor := 1000
	for i := 1; i <= N; i++ {
		soma += valor
		valor -= 3
	}
	fmt.Println("Resultado da série =", soma)
}
