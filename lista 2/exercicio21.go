package main

import "fmt"

func main() {
	var id int
	var n1, n2, n3, mex float64
	fmt.Scan(&id, &n1, &n2, &n3, &mex)
	media := (n1 + n2*2 + n3*3 + mex) / 7
	conceito := ""
	if media >= 9 {
		conceito = "A"
	} else if media >= 7.5 {
		conceito = "B"
	} else if media >= 6 {
		conceito = "C"
	} else if media >= 4 {
		conceito = "D"
	} else {
		conceito = "E"
	}
	fmt.Println("Aluno:", id, "Média:", media, "Conceito:", conceito)
	if conceito == "A" || conceito == "B" || conceito == "C" {
		fmt.Println("APROVADO")
	} else {
		fmt.Println("REPROVADO")
	}
}
