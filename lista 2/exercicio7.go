package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("Digite três números distintos: ")
	fmt.Scan(&a, &b, &c)

	menor, inter, maior := a, b, c

	if menor > b {
		menor, b = b, menor
	}
	if menor > c {
		menor, c = c, menor
	}
	if b > c {
		b, c = c, b
	}

	inter, maior = b, c

	fmt.Println("Ordem crescente:", menor, inter, maior)
}
