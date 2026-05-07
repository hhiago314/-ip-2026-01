package main

import "fmt"

func main() {
	var notas [15]int
	freqAbs := make(map[int]int)

	for i := 0; i < 15; i++ {
		fmt.Printf("Digite a nota %d (0-10): ", i+1)
		fmt.Scan(&notas[i])
		freqAbs[notas[i]]++
	}

	fmt.Println("Nota | Freq. Absoluta | Freq. Relativa")
	for nota := 0; nota <= 10; nota++ {
		abs := freqAbs[nota]
		rel := float64(abs) / 15.0
		fmt.Printf("%4d | %14d | %13.2f\n", nota, abs, rel)
	}
}
