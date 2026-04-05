package main

import "fmt"

func main() {
	var num int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&num)

	if num%5 == 0 {
		fmt.Println("O número é divisível por 5.")
	} else {
		fmt.Println("O número não é divisível por 5.")
	}
}
