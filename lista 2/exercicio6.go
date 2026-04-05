package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Digite o valor de A: ")
	fmt.Scan(&a)
	fmt.Print("Digite o valor de B: ")
	fmt.Scan(&b)

	if b != 0 && a%b == 0 {
		fmt.Println("A é divisível por B.")
	} else {
		fmt.Println("A não é divisível por B.")
	}
}
