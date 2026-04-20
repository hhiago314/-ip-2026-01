package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite um número decimal: ")
	fmt.Scan(&n)

	bin := ""
	for n > 0 {
		if n%2 == 0 {
			bin = "0" + bin
		} else {
			bin = "1" + bin
		}
		n /= 2
	}
	fmt.Println("Binário =", bin)
}
