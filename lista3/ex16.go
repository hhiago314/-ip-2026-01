package main

import "fmt"

func main() {
	var n, a1, a2 int
	fmt.Print("Digite N (>=3): ")
	fmt.Scan(&n)
	fmt.Print("Digite os dois primeiros termos: ")
	fmt.Scan(&a1, &a2)

	fmt.Print(a1, " ", a2, " ")
	for i := 3; i <= n; i++ {
		var ai int
		if i%2 == 1 {
			ai = a1 + a2
		} else {
			ai = a1 - a2
		}
		fmt.Print(ai, " ")
		a1, a2 = a2, ai
	}
	fmt.Println()
}
