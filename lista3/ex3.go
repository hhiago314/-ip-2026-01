package main

import "fmt"

func main() {
	var salarioCarlos float64

	fmt.Print("Digite o salário de Carlos: ")
	fmt.Scan(&salarioCarlos)

	salarioJoao := salarioCarlos / 3

	investCarlos := salarioCarlos
	investJoao := salarioJoao

	meses := 0

	for investJoao < investCarlos {
		investCarlos *= 1.02
		investJoao *= 1.05
		meses++
	}

	fmt.Println("Meses necessários:", meses)
}
