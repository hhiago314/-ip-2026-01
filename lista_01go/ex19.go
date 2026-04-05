package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n <= 1 {
		fmt.Println("Numero invalido!")
		return
	}
	soma := 0.0
	for i := 1; i <= n; i++ {
		soma += 1.0 / float64(i)
	}
	fmt.Printf("%.6f\n", soma)
}
