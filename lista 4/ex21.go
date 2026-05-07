package main

import "fmt"

func main() {
	var v [10]float64
	var codigo int

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite v[%d]: ", i)
		fmt.Scan(&v[i])
	}

	for {
		fmt.Print("Digite código (0=termina, 1=ordem direta, 2=ordem inversa): ")
		fmt.Scan(&codigo)

		if codigo == 0 {
			break
		} else if codigo == 1 {
			fmt.Println("Ordem direta:", v)
		} else if codigo == 2 {
			fmt.Print("Ordem inversa: ")
			for i := 9; i >= 0; i-- {
				fmt.Print(v[i], " ")
			}
			fmt.Println()
		}
	}
}
