package main

import "fmt"

func main() {
	type Empregado struct {
		codigo int
		meses  int
	}

	var empregados []Empregado

	for {
		var cod, meses int
		fmt.Print("Digite código e meses (0 0 para parar): ")
		fmt.Scan(&cod, &meses)
		if cod == 0 && meses == 0 {
			break
		}
		empregados = append(empregados, Empregado{cod, meses})
	}

	// ordenar por meses decrescente
	for i := 0; i < len(empregados); i++ {
		for j := i + 1; j < len(empregados); j++ {
			if empregados[i].meses < empregados[j].meses {
				empregados[i], empregados[j] = empregados[j], empregados[i]
			}
		}
	}

	fmt.Println("Três mais recentes:")
	for i := 0; i < 3 && i < len(empregados); i++ {
		fmt.Printf("Empregado %d - %d meses\n", empregados[i].codigo, empregados[i].meses)
	}
}
