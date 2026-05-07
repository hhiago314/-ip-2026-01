package main

import "fmt"

func main() {
	var alturas [10]float64
	soma := 0.0

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite a altura do atleta %d: ", i+1)
		fmt.Scan(&alturas[i])
		soma += alturas[i]
	}

	media := soma / 10
	fmt.Println("Alturas maiores que a média:")
	for _, h := range alturas {
		if h > media {
			fmt.Println(h)
		}
	}
}
