package main

import "fmt"

func main() {
	preco := 6.0
	maxLucro := -1.0
	melhorPreco := 0.0
	melhorIngressos := 0

	for preco >= 1.0 {
		ingressos := 130 + int((6.0-preco)/0.6*30)
		receita := preco * float64(ingressos)
		lucro := receita - 300.0

		fmt.Printf("Preço=%.2f Ingressos=%d Lucro=%.2f\n", preco, ingressos, lucro)

		if lucro > maxLucro {
			maxLucro = lucro
			melhorPreco = preco
			melhorIngressos = ingressos
		}

		preco -= 0.6
	}

	fmt.Printf("\nLucro máximo=%.2f com preço=%.2f e ingressos=%d\n",
		maxLucro, melhorPreco, melhorIngressos)
}
