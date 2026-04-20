package main

import "fmt"

func main() {
	soma := 0
	for i := 1; i <= 37; i++ {
		soma += i*(i+1) - i
	}
	fmt.Println("Resultado da soma =", soma)
}
