package main

import (
	"fmt"
)

func main() {
	var conta int
	var consumo float64
	var tipo rune

	fmt.Scan(&conta, &consumo, &tipo)

	var valor float64

	switch tipo {
	case 'R':
		valor = 5.0 + (0.05 * consumo)
	case 'C':
		if consumo <= 80 {
			valor = 500.0
		} else {
			valor = 500.0 + (0.25 * (consumo - 80))
		}
	case 'I':
		if consumo <= 100 {
			valor = 800.0
		} else {
			valor = 800.0 + (0.04 * (consumo - 100))
		}
	default:
		fmt.Println("TIPO INVALIDO")
		return
	}

	fmt.Printf("CONTA = %d\n", conta)
	fmt.Printf("VALOR DA CONTA = %.2f\n", valor)
}
