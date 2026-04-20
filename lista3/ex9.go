package main

import "fmt"

func main() {
	var N int
	fmt.Print("Digite o número de alunos: ")
	fmt.Scan(&N)

	aprovados, exames, reprovados := 0, 0, 0
	somaClasse := 0.0

	for i := 1; i <= N; i++ {
		var n1, n2 float64
		fmt.Printf("Digite as duas notas do aluno %d: ", i)
		fmt.Scan(&n1, &n2)

		media := (n1 + n2) / 2
		somaClasse += media

		fmt.Printf("Média do aluno %d = %.2f -> ", i, media)
		if media < 3 {
			fmt.Println("Reprovado")
			reprovados++
		} else if media < 7 {
			fmt.Println("Exame")
			exames++
		} else {
			fmt.Println("Aprovado")
			aprovados++
		}
	}

	fmt.Println("\nResumo da turma:")
	fmt.Println("Total de aprovados:", aprovados)
	fmt.Println("Total de exames:", exames)
	fmt.Println("Total de reprovados:", reprovados)
	fmt.Printf("Média da classe: %.2f\n", somaClasse/float64(N))
}
