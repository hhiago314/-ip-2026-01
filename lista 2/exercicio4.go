package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64
	fmt.Print("Digite um número: ")
	fmt.Scan(&num)

	if num >= 0 {
		fmt.Println("Raiz quadrada:", math.Sqrt(num))
	} else {
		fmt.Println("Quadrado:", num*num)
	}
}
