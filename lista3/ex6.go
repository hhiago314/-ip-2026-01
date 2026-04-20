package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite um número: ")
	fmt.Scan(&n)

	triangular := false
	for i := 1; i*(i+1)*(i+2) <= n; i++ {
		if i*(i+1)*(i+2) == n {
			triangular = true
			break
		}
	}

	if triangular {
		fmt.Println(n, "é triangular")
	} else {
		fmt.Println(n, "não é triangular")
	}
}
