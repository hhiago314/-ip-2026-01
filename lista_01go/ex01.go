package main

import (
	"fmt"
)

func main() {
	var n1, n2, n3 float64

	fmt.Scanf("%f %f %f", &n1, &n2, &n3)

	media := (n1 + n2 + n3) / 3.0

	fmt.Printf("MEDIA = %.2f\n", media)

	if media >= 6.0 {
		fmt.Println("APROVADO")
	} else {
		fmt.Println("REPROVADO")
	}
}
