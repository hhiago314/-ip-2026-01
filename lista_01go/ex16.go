package main

import "fmt"

func main() {
	var salario float64
	fmt.Scan(&salario)
	if salario <= 300 {
		salario *= 1.5
	} else {
		salario *= 1.3
	}
	fmt.Printf("SALARIO COM REAJUSTE = %.2f\n", salario)
}
