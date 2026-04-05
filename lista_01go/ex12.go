package main

import "fmt"

func main() {
	var horas int
	fmt.Scan(&horas)

	blocos := horas / 3
	resto := horas % 3

	valor := float64(blocos*10 + resto*5)

	fmt.Printf("O VALOR A PAGAR E = %.2f\n", valor)
}
