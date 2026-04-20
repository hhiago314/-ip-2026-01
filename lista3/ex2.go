package main

import "fmt"

func main() {
	soma := 0
	qtd := 0
	for i := 50; i <= 70; i++ {
		if i%2 == 0 {
			soma += i
			qtd++
		}
	}
	media := float64(soma) / float64(qtd)
	fmt.Printf("Soma = %d, Média = %.2f\n", soma, media)
}
