package main

import "fmt"

func main() {
	var tipo string
	var consumo float64
	fmt.Scan(&tipo, &consumo)
	var conta float64
	switch tipo {
	case "residencial":
		conta = 5 + 0.05*consumo
	case "comercial":
		if consumo <= 80 {
			conta = 500
		} else {
			conta = 500 + 0.25*(consumo-80)
		}
	case "industrial":
		if consumo <= 100 {
			conta = 800
		} else {
			conta = 800 + 0.04*(consumo-100)
		}
	default:
		fmt.Println("Tipo inválido")
		return
	}
	fmt.Println("Conta:", conta)
}
