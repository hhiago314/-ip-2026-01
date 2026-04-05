package main

import (
	"fmt"
)

func main() {
	var fahrenheit, polegadas float64

	fmt.Scan(&fahrenheit)
	fmt.Scan(&polegadas)

	celsius := (5*fahrenheit - 160) / 9.0

	milimetros := polegadas * 25.4

	fmt.Printf("O VALOR EM CELSIUS = %.2f\n", celsius)
	fmt.Printf("A QUANTIDADE DE CHUVA E = %.2f\n", milimetros)
}
