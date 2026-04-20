package main

import (
	"fmt"
)

func reverseArray(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	var n int
	fmt.Print("Digite o tamanho do array: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Digite o valor %d: ", i+1)
		fmt.Scan(&arr[i])
	}

	reverseArray(arr)
	fmt.Println("Array invertido:", arr)
}
