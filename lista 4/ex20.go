package main

import "fmt"

func main() {
	var jogadas [20]int
	freq := make(map[int]int)

	for i := 0; i < 20; i++ {
		fmt.Printf("Digite jogada %d (1-6): ", i+1)
		fmt.Scan(&jogadas[i])
		freq[jogadas[i]]++
	}

	fmt.Println("Frequência dos números sorteados:")
	for num := 1; num <= 6; num++ {
		fmt.Printf("%d apareceu %d vezes\n", num, freq[num])
	}
}
