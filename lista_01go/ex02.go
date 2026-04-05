package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		var total int
		var pPopular, pGeral, pArquibancada, pCadeiras float64

		fmt.Scan(&total, &pPopular, &pGeral, &pArquibancada, &pCadeiras)

		popular := float64(total) * pPopular / 100.0
		geral := float64(total) * pGeral / 100.0
		arquibancada := float64(total) * pArquibancada / 100.0
		cadeiras := float64(total) * pCadeiras / 100.0

		renda := popular*1.0 + geral*5.0 + arquibancada*10.0 + cadeiras*20.0

		fmt.Printf("A RENDA DO JOGO N. %d E = %.2f\n", i, renda)
	}
}
