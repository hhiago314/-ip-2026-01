package main

import "fmt"

func main() {
	var idade int
	fmt.Scan(&idade)
	switch {
	case idade >= 0 && idade <= 2:
		fmt.Println("Recém-nascido")
	case idade <= 11:
		fmt.Println("Criança")
	case idade <= 19:
		fmt.Println("Adolescente")
	case idade <= 55:
		fmt.Println("Adulto")
	default:
		fmt.Println("Idoso")
	}
}
