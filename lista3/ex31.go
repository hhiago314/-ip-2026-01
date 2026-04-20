package main

import "fmt"

func main() {
	graos := uint64(1)
	soma := uint64(1)
	for i := 2; i <= 64; i++ {
		graos *= 2
		soma += graos
	}
	fmt.Println("Total de grãos =", soma)
}
