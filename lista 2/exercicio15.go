package main

import "fmt"

func main() {
	var d, m, a int
	fmt.Scan(&d, &m, &a)
	meses := []string{"Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho", "Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro"}
	if m >= 1 && m <= 12 {
		fmt.Printf("%d de %s de %d\n", d, meses[m-1], a)
	} else {
		fmt.Println("Mês inválido")
	}
}
