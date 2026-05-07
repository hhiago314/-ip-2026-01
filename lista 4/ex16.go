package main

import "fmt"

func main() {
	var idades [50]int
	freq := make(map[int]int)

	for i := 0; i < 50; i++ {
		fmt.Printf("Digite idade %d: ", i+1)
		fmt.Scan(&idades[i])
		freq[idades[i]]++
	}

	moda := idades[0]
	maxFreq := 0
	for idade, qtd := range freq {
		if qtd > maxFreq {
			maxFreq = qtd
			moda = idade
		}
	}

	fmt.Printf("Moda das idades: %d (aparece %d vezes)\n", moda, maxFreq)
}
