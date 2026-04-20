package main

import "fmt"

func main() {
	var N int
	fmt.Print("Digite N: ")
	fmt.Scan(&N)

	soma := 0
	for i := 1; i <= N; i++ {
		soma += i
	}
	fmt.Println("Soma =", soma)
}
