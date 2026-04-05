package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&num)

	if num > 0 {
		fmt.Println("Positivo")
	} else if num < 0 {
		fmt.Println("Negativo")
	} else {
		fmt.Println("Nulo")
	}
}
