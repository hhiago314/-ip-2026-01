package main

import "fmt"

func main() {
	var destino, retorno int
	fmt.Scan(&destino, &retorno)
	tabela := map[int][2]float64{
		1: {500, 900},
		2: {350, 650},
		3: {350, 600},
		4: {300, 550},
	}
	if val, ok := tabela[destino]; ok {
		if retorno == 1 {
			fmt.Println("Preço:", val[1])
		} else if retorno == 2 {
			fmt.Println("Preço:", val[0])
		} else {
			fmt.Println("Retorno inválido")
		}
	} else {
		fmt.Println("Destino inválido")
	}
}
