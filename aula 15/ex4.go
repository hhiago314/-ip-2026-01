package main

import (
	"fmt"
)

func decimalToBinary(n int) string {
	if n == 0 {
		return "0"
	}
	binary := ""
	for n > 0 {
		if n%2 == 0 {
			binary = "0" + binary
		} else {
			binary = "1" + binary
		}
		n /= 2
	}
	return binary
}

func main() {
	var n int
	fmt.Print("Digite um número decimal: ")
	fmt.Scan(&n)

	fmt.Printf("Binário: %s\n", decimalToBinary(n))
}
