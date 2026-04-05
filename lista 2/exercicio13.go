package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	if num >= 100 && num <= 999 {
		dezenas := (num / 10) % 10
		fmt.Println("Dezena:", dezenas)
	} else {
		fmt.Println("Número não tem 3 dígitos")
	}
}
