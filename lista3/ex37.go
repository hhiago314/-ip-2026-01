package main

import "fmt"

func main() {
	var n string
	fmt.Print("Digite um número em octal: ")
	fmt.Scan(&n)

	decimal := 0
	base := 1
	for i := len(n) - 1; i >= 0; i-- {
		decimal += int(n[i]-'0') * base
		base *= 8
	}
	fmt.Println("Decimal =", decimal)
}
