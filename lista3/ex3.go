package main

import "fmt"

func main() {
	var salarioCarlos float64
	fmt.Print("Digite o salário de Carlos: ")
	fmt.Scan(&salarioCarlos)

	salarioJoao := salarioCarlos / 3
	meses := 0

	for salarioJoao <= salarioCarlos {
		salarioCarlos *= 1.02 // 2% ao mês
		salarioJoao *= 1.05   // 5% ao mês
		meses++
	}

	fmt.Printf("João ultrapassa Carlos em %d meses.\n", meses)
}
