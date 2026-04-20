package main

import "fmt"

func main() {
	var N int
	fmt.Print("Digite N (>=3): ")
	fmt.Scan(&N)

	a, b := 0, 1
	fmt.Print(a, " ", b, " ")
	for i := 3; i <= N; i++ {
		c := a + b
		fmt.Print(c, " ")
		a, b = b, c
	}
	fmt.Println()
}
