package main

import (
	"fmt"
)

func main() {
	var salarioMinimo, consumo float64

	fmt.Scan(&salarioMinimo, &consumo)

	valorKw := (salarioMinimo * 0.70) / 100.0

	custoConsumo := valorKw * consumo

	custoDesconto := custoConsumo * 0.90

	fmt.Printf("Custo por kW: R$ %.2f\n", valorKw)
	fmt.Printf("Custo do consumo: R$ %.2f\n", custoConsumo)
	fmt.Printf("Custo com desconto: R$ %.2f\n", custoDesconto)
}
