package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite um número decimal: ")
	fmt.Scan(&n)

	hex := ""
	letras := "0123456789ABCDEF"
	for n > 0 {
		hex = string(letras[n%16]) + hex
		n /= 16
	}
	fmt.Println("Hexadecimal =", hex)
}
