package main

import "fmt"

func main() {
	var idade int
	var altura, peso float64
	var qtd50, qtd1020 int
	var somaAlturas float64
	var qtdTotal, qtdPeso40 int

	for {
		fmt.Print("Digite idade (0 para sair): ")
		fmt.Scan(&idade)
		if idade == 0 {
			break
		}
		fmt.Print("Digite altura: ")
		fmt.Scan(&altura)
		fmt.Print("Digite peso: ")
		fmt.Scan(&peso)

		qtdTotal++
		if idade > 50 {
			qtd50++
		}
		if idade >= 10 && idade <= 20 {
			somaAlturas += altura
			qtd1020++
		}
		if peso < 40 {
			qtdPeso40++
		}
	}

	fmt.Println("Qtd > 50 anos:", qtd50)
	if qtd1020 > 0 {
		fmt.Println("Média alturas 10–20:", somaAlturas/float64(qtd1020))
	}
	fmt.Printf("%% com peso < 40kg: %.2f%%\n", float64(qtdPeso40)/float64(qtdTotal)*100)
}
