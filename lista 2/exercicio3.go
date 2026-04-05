package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Print("Digite dois números inteiros: ")
	fmt.Scan(&a, &b)

	soma := a + b
	if soma > 20 {
		fmt.Println("Resultado:", soma+8)
	} else {
		fmt.Println("Resultado:", soma-5)
	}
}
