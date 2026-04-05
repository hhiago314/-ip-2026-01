package main

import "fmt"

func main() {
	var num int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&num)

	if num > 20 && num < 90 {
		fmt.Println("O número está entre 20 e 90.")
	} else {
		fmt.Println("O número NÃO está entre 20 e 90.")
	}
}
