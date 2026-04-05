package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	delta := b*b - 4*a*c
	if delta < 0 {
		fmt.Println("Raízes imaginárias")
	} else if delta == 0 {
		x := -b / (2 * a)
		fmt.Println("Raiz única:", x)
	} else {
		x1 := (-b + math.Sqrt(delta)) / (2 * a)
		x2 := (-b - math.Sqrt(delta)) / (2 * a)
		fmt.Println("Raízes distintas:", x1, x2)
	}
}
