package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Print("Digite N1 e N2: ")
	fmt.Scan(&n1, &n2)

	quociente := 0
	for n1 >= n2 {
		n1 -= n2
		quociente++
	}
	fmt.Printf("Quociente = %d, Resto = %d\n", quociente, n1)
}
