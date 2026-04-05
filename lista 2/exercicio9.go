package main

import "fmt"

func main() {
	var compra float64
	fmt.Scan(&compra)
	if compra < 0 {
		fmt.Println("Valor inválido")
		return
	}
	var venda float64
	switch {
	case compra < 10:
		venda = compra * 1.7
	case compra < 30:
		venda = compra * 1.5
	case compra < 50:
		venda = compra * 1.4
	default:
		venda = compra * 1.3
	}
	fmt.Println("Valor da venda:", venda)
}
