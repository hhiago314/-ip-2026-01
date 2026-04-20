package main

import "fmt"

func main() {
	for {
		var n int
		fmt.Print("Digite um número (<=0 para sair): ")
		fmt.Scan(&n)

		if n <= 0 {
			break
		}

		quadrado := false
		for i := 1; i*i <= n; i++ {
			if i*i == n {
				quadrado = true
				break
			}
		}

		if quadrado {
			fmt.Println(n, "é quadrado perfeito")
		} else {
			fmt.Println(n, "não é quadrado perfeito")
		}
	}
}
