package main

import "fmt"

func main() {
	var preco float64
	var ar, pintura, vidro, direcao int
	fmt.Scan(&preco, &ar, &pintura, &vidro, &direcao)
	if ar == 1 {
		preco += 1750
	}
	if pintura == 1 {
		preco += 800
	}
	if vidro == 1 {
		preco += 1200
	}
	if direcao == 1 {
		preco += 2000
	}
	fmt.Println("Preço final:", preco)
}
