package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scan(&n)

	if n < 0 {
		fmt.Println("Número inválido!")
		return
	}

	fat := 1
	for i := 2; i <= n; i++ {
		fat *= i
	}
	fmt.Printf("%d! = %d\n", n, fat)
}
