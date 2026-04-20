package main

import "fmt"

func main() {
	var cpf [11]int
	fmt.Println("Digite os 11 dígitos do CPF:")
	for i := 0; i < 11; i++ {
		fmt.Scan(&cpf[i])
	}

	// Primeiro dígito verificador
	soma := 0
	for i := 0; i < 9; i++ {
		soma += cpf[i] * (10 - i)
	}
	resto := soma % 11
	d1 := 0
	if resto >= 2 {
		d1 = 11 - resto
	}

	// Segundo dígito verificador
	soma = 0
	for i := 0; i < 10; i++ {
		soma += cpf[i] * (11 - i)
	}
	resto = soma % 11
	d2 := 0
	if resto >= 2 {
		d2 = 11 - resto
	}

	if d1 == cpf[9] && d2 == cpf[10] {
		fmt.Println("CPF válido")
	} else {
		fmt.Println("CPF inválido")
	}
}
