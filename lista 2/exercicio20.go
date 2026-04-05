package main

import "fmt"

func main() {
	var preco float64
	var codigo int
	fmt.Scan(&preco, &codigo)
	switch codigo {
	case 1:
		fmt.Println("Valor final:", preco*0.9)
	case 2:
		fmt.Println("Valor final:", preco*0.95)
	case 3:
		fmt.Println("Valor final:", preco)
	case 4:
		fmt.Println("Valor final:", preco*1.1)
	default:
		fmt.Println("Código inválido")
	}
}
