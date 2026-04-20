package main

import "fmt"

func main() {
	var n int
	soma, qtd, somaPares, qtdPares := 0, 0, 0, 0
	maior, menor := -1<<31, 1<<31-1
	qtdImpares := 0

	for {
		fmt.Print("Digite um número (30000 para sair): ")
		fmt.Scan(&n)
		if n == 30000 {
			break
		}
		soma += n
		qtd++
		if n > maior {
			maior = n
		}
		if n < menor {
			menor = n
		}
		if n%2 == 0 {
			somaPares += n
			qtdPares++
		} else {
			qtdImpares++
		}
	}

	fmt.Println("Soma:", soma)
	fmt.Println("Qtd:", qtd)
	fmt.Println("Média:", float64(soma)/float64(qtd))
	fmt.Println("Maior:", maior)
	fmt.Println("Menor:", menor)
	if qtdPares > 0 {
		fmt.Println("Média pares:", float64(somaPares)/float64(qtdPares))
	}
	fmt.Printf("%% ímpares: %.2f%%\n", float64(qtdImpares)/float64(qtd)*100)
}
