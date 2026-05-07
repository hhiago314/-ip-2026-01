package main

import "fmt"

func main() {
	var nums [10]int
	var divis [5]int

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite nums[%d]: ", i)
		fmt.Scan(&nums[i])
	}
	for i := 0; i < 5; i++ {
		fmt.Printf("Digite divis[%d]: ", i)
		fmt.Scan(&divis[i])
	}

	for i, num := range nums {
		fmt.Printf("Número %d:\n", num)
		for j, d := range divis {
			if num%d == 0 {
				fmt.Printf("  Divisível por %d na posição %d\n", d, j)
			}
		}
	}
}
