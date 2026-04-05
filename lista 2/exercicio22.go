package main

import "fmt"

func main() {
	var matricula, horas int
	fmt.Scan(&matricula, &horas)
	salarioMin := 788.0
	extra := float64(horas) * 10
	bruto := 3*salarioMin + extra
	desconto := 0.0
	if bruto > 1500 {
		desconto += bruto * 0.12
	}
	if bruto > 2000 {
		desconto += bruto * 0.20
	}
	liquido := bruto - desconto
	fmt.Println("Matrícula:", matricula, "Bruto:", bruto, "Líquido:", liquido)
}
