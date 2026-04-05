package main

import (
	"fmt"
	"math"
)

func main() {
	var opcao int
	var raio, altura float64
	fmt.Scan(&opcao)
	switch opcao {
	case 1: // Cone
		fmt.Scan(&raio, &altura)
		volume := (math.Pi * raio * raio * altura) / 3
		area := math.Pi * raio * (raio + math.Sqrt(raio*raio+altura*altura))
		fmt.Println("Volume:", volume, "Área:", area)
	case 2: // Cilindro
		fmt.Scan(&raio, &altura)
		volume := math.Pi * raio * raio * altura
		area := 2 * math.Pi * raio * (raio + altura)
		fmt.Println("Volume:", volume, "Área:", area)
	case 3: // Esfera
		fmt.Scan(&raio)
		volume := (4.0 / 3.0) * math.Pi * math.Pow(raio, 3)
		area := 4 * math.Pi * raio * raio
		fmt.Println("Volume:", volume, "Área:", area)
	default:
		fmt.Println("Opção inválida")
	}
}
